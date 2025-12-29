package parser

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/promonkeyli/goas/pkg/model"
	"golang.org/x/tools/go/packages"
)

// resolveTypeSchema 解析类型字符串并返回 Schema
// 支持: 基本类型, 包路径.类型名, 泛型类型 (如 Response[User])
func (p *Processor) resolveTypeSchema(pkg *packages.Package, typeName string) *model.Schema {
	if typeName == "" {
		return nil
	}

	// 处理泛型类型 (如 Response[User], Page[model.Item])
	if genericBase, typeArgs := parseGenericType(typeName); genericBase != "" {
		return p.resolveGenericSchema(pkg, genericBase, typeArgs)
	}

	// 处理基本类型
	if schema := p.primitiveTypeToSchema(typeName); schema != nil {
		if typeStr, ok := schema.Type.(string); ok && typeStr != "" {
			return schema
		}
	}

	// 处理复合类型引用
	return p.resolveStructSchema(pkg, typeName)
}

// parseGenericType 解析泛型类型名称
// 输入: "Response[User]" 或 "Page[model.Item, int]"
// 输出: ("Response", ["User"]) 或 ("Page", ["model.Item", "int"])
func parseGenericType(typeName string) (base string, typeArgs []string) {
	start := strings.Index(typeName, "[")
	end := strings.LastIndex(typeName, "]")

	if start == -1 || end == -1 || start >= end {
		return "", nil
	}

	base = typeName[:start]
	argsStr := typeName[start+1 : end]

	// 分割类型参数
	var args []string
	depth := 0
	current := strings.Builder{}

	for _, r := range argsStr {
		switch r {
		case '[':
			depth++
			current.WriteRune(r)
		case ']':
			depth--
			current.WriteRune(r)
		case ',':
			if depth == 0 {
				args = append(args, strings.TrimSpace(current.String()))
				current.Reset()
			} else {
				current.WriteRune(r)
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		args = append(args, strings.TrimSpace(current.String()))
	}

	return base, args
}

// resolveGenericSchema 解析泛型类型的 Schema
func (p *Processor) resolveGenericSchema(pkg *packages.Package, baseType string, typeArgs []string) *model.Schema {
	// 生成唯一的 Schema 名称 (如 ResponseOfUser)
	schemaName := generateGenericSchemaName(baseType, typeArgs)

	// 检查缓存
	if ref, ok := p.GeneratedSchemas[schemaName]; ok {
		return &model.Schema{Ref: ref}
	}

	// 预先添加到缓存，防止循环引用
	refPath := "#/components/schemas/" + schemaName
	p.GeneratedSchemas[schemaName] = refPath

	// 解析包路径和类型名
	pkgPath, shortName := parseTypePath(pkg, baseType)

	// 查找类型定义
	targetPkg := p.findPackage(pkgPath)
	if targetPkg == nil || targetPkg.Types == nil {
		return &model.Schema{Ref: refPath}
	}

	// 查找类型对象
	obj := targetPkg.Types.Scope().Lookup(shortName)
	if obj == nil {
		return &model.Schema{Ref: refPath}
	}

	// 获取底层结构体
	namedType, ok := obj.Type().(*types.Named)
	if !ok {
		return &model.Schema{Ref: refPath}
	}

	underlying, ok := namedType.Underlying().(*types.Struct)
	if !ok {
		return &model.Schema{Ref: refPath}
	}

	// 创建实例化的 Schema
	schema := &model.Schema{
		Type:       "object",
		Title:      schemaName,
		Properties: make(map[string]*model.Schema),
	}

	// 遍历结构体字段
	for i := 0; i < underlying.NumFields(); i++ {
		field := underlying.Field(i)
		tag := underlying.Tag(i)

		if !field.Exported() {
			continue
		}

		jsonName, omitempty := parseJSONTag(tag)
		if jsonName == "-" {
			continue
		}
		if jsonName == "" {
			jsonName = field.Name()
		}

		// 使用带类型参数替换的类型解析 (使用 targetPkg 确保类型在正确的包上下文中解析)
		fieldSchema := p.typeToSchemaWithSubstitution(targetPkg, field.Type(), typeArgs)

		// 添加描述
		if desc := parseDescTag(tag); desc != "" {
			fieldSchema.Description = desc
		}

		schema.Properties[jsonName] = fieldSchema

		if !omitempty {
			schema.Required = append(schema.Required, jsonName)
		}
	}

	// 添加到 Components
	p.ensureComponents()
	if p.OpenAPI.Components.Schemas == nil {
		p.OpenAPI.Components.Schemas = make(map[string]*model.Schema)
	}
	p.OpenAPI.Components.Schemas[schemaName] = schema

	return &model.Schema{Ref: refPath}
}

// generateGenericSchemaName 生成泛型 Schema 名称
// 输入: ("Response", ["User"]) -> "ResponseOfUser"
// 输入: ("Page", ["model.Item"]) -> "PageOfItem"
func generateGenericSchemaName(base string, typeArgs []string) string {
	var parts []string
	for _, arg := range typeArgs {
		// 取类型的最后一部分 (model.User -> User)
		if idx := strings.LastIndex(arg, "."); idx != -1 {
			parts = append(parts, arg[idx+1:])
		} else {
			parts = append(parts, arg)
		}
	}
	return base + "Of" + strings.Join(parts, "And")
}

// instantiateGenericSchema 实例化泛型 Schema
func (p *Processor) instantiateGenericSchema(pkg *packages.Package, base *model.Schema, typeArgs []string) *model.Schema {
	// 深拷贝基础 Schema
	schema := &model.Schema{
		Type:        base.Type,
		Description: base.Description,
		Required:    base.Required,
	}

	if base.Properties != nil {
		schema.Properties = make(map[string]*model.Schema)
		for name, prop := range base.Properties {
			// 检查是否是类型参数占位符
			if prop.Ref != "" && isTypeParamPlaceholder(prop.Ref) {
				// 用实际类型替换
				if len(typeArgs) > 0 {
					schema.Properties[name] = p.resolveTypeSchema(pkg, typeArgs[0])
				}
			} else {
				schema.Properties[name] = prop
			}
		}
	}

	return schema
}

// typeToSchemaWithSubstitution 将 Go 类型转换为 Schema，同时替换类型参数
func (p *Processor) typeToSchemaWithSubstitution(pkg *packages.Package, t types.Type, typeArgs []string) *model.Schema {
	switch typ := t.(type) {
	case *types.TypeParam:
		// 类型参数，用实际类型替换
		paramIndex := typ.Index()
		if paramIndex < len(typeArgs) {
			return p.resolveTypeSchema(pkg, typeArgs[paramIndex])
		}
		return &model.Schema{Type: "object"}

	case *types.Pointer:
		return p.typeToSchemaWithSubstitution(pkg, typ.Elem(), typeArgs)

	case *types.Slice:
		return &model.Schema{
			Type:  "array",
			Items: p.typeToSchemaWithSubstitution(pkg, typ.Elem(), typeArgs),
		}

	case *types.Array:
		return &model.Schema{
			Type:     "array",
			Items:    p.typeToSchemaWithSubstitution(pkg, typ.Elem(), typeArgs),
			MaxItems: int(typ.Len()),
			MinItems: int(typ.Len()),
		}

	case *types.Map:
		return &model.Schema{
			Type:                 "object",
			AdditionalProperties: p.typeToSchemaWithSubstitution(pkg, typ.Elem(), typeArgs),
		}

	case *types.Named:
		// 递归解析命名类型
		obj := typ.Obj()
		pkgPath := ""
		if obj.Pkg() != nil {
			pkgPath = obj.Pkg().Path()
		}
		fullName := pkgPath + "." + obj.Name()

		// 检查缓存
		if ref, ok := p.GeneratedSchemas[fullName]; ok {
			return &model.Schema{Ref: ref}
		}

		// 处理泛型类型
		if typeParams := typ.TypeParams(); typeParams != nil && typeParams.Len() > 0 {
			return p.handleGenericNamedType(typ)
		}

		// 解析底层类型
		return p.typeToSchemaWithSubstitution(pkg, typ.Underlying(), typeArgs)

	case *types.Struct:
		return p.structToSchema(nil, typ, "")

	case *types.Interface:
		return &model.Schema{Type: "object"}

	case *types.Basic:
		return p.basicTypeToSchema(typ)

	default:
		return p.typeToSchema(t)
	}
}

// isTypeParamPlaceholder 检查是否是类型参数占位符
func isTypeParamPlaceholder(ref string) bool {
	// 类型参数通常是单个大写字母 (T, K, V 等)
	if strings.HasPrefix(ref, "#/components/schemas/") {
		name := strings.TrimPrefix(ref, "#/components/schemas/")
		return len(name) == 1 && name[0] >= 'A' && name[0] <= 'Z'
	}
	return false
}

// resolveStructSchema 解析结构体类型的 Schema
func (p *Processor) resolveStructSchema(pkg *packages.Package, typeName string) *model.Schema {
	// 处理 [] 切片类型
	if strings.HasPrefix(typeName, "[]") {
		itemType := strings.TrimPrefix(typeName, "[]")
		return &model.Schema{
			Type:  "array",
			Items: p.resolveTypeSchema(pkg, itemType),
		}
	}

	// 处理 map 类型
	if strings.HasPrefix(typeName, "map[") {
		return &model.Schema{
			Type:                 "object",
			AdditionalProperties: true,
		}
	}

	// 解析包路径和类型名
	pkgPath, shortName := parseTypePath(pkg, typeName)

	// 检查缓存
	fullName := pkgPath + "." + shortName
	if ref, ok := p.GeneratedSchemas[fullName]; ok {
		return &model.Schema{Ref: ref}
	}

	// 查找类型定义
	// 如果类型在当前包中，直接使用当前包，避免 findPackage 查找失败
	var targetPkg *packages.Package
	if pkg != nil && pkgPath == pkg.PkgPath {
		targetPkg = pkg
	} else {
		targetPkg = p.findPackage(pkgPath)
	}
	if targetPkg == nil || targetPkg.Types == nil {
		// 无法找到包，返回空 object
		return &model.Schema{Type: "object"}
	}

	// 查找类型对象
	obj := targetPkg.Types.Scope().Lookup(shortName)
	if obj == nil {
		return &model.Schema{Type: "object"}
	}

	typeDef, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		// 可能是类型别名或基本类型
		return p.typeToSchema(obj.Type())
	}

	// 预先添加到缓存，防止循环引用
	refPath := "#/components/schemas/" + shortName
	p.GeneratedSchemas[fullName] = refPath

	// 解析结构体字段
	schema := p.structToSchema(targetPkg, typeDef, shortName)

	// 添加到 Components
	p.ensureComponents()
	if p.OpenAPI.Components.Schemas == nil {
		p.OpenAPI.Components.Schemas = make(map[string]*model.Schema)
	}
	p.OpenAPI.Components.Schemas[shortName] = schema

	return &model.Schema{Ref: refPath}
}

// parseTypePath 解析类型路径
// 输入: "model.User" 或 "User"
// 输出: (包路径, 类型名)
func parseTypePath(pkg *packages.Package, typeName string) (pkgPath, shortName string) {
	if idx := strings.LastIndex(typeName, "."); idx != -1 {
		// 有包前缀
		pkgAlias := typeName[:idx]
		shortName = typeName[idx+1:]

		// 在当前文件的 imports 中查找实际包路径
		for importPath := range pkg.Imports {
			if strings.HasSuffix(importPath, "/"+pkgAlias) || strings.HasSuffix(importPath, pkgAlias) {
				pkgPath = importPath
				break
			}
		}

		if pkgPath == "" {
			pkgPath = pkg.PkgPath // 假设是同包
		}
	} else {
		// 没有包前缀，使用当前包
		pkgPath = pkg.PkgPath
		shortName = typeName
	}

	return pkgPath, shortName
}

// findPackage 在已加载的包中查找指定路径的包
func (p *Processor) findPackage(pkgPath string) *packages.Package {
	if pkg, ok := p.PackagesMap[pkgPath]; ok {
		return pkg
	}

	// 尝试模糊匹配
	for path, pkg := range p.PackagesMap {
		if strings.HasSuffix(path, "/"+pkgPath) || strings.HasSuffix(path, pkgPath) {
			return pkg
		}
	}

	return nil
}

// structToSchema 将 Go 结构体转换为 Schema
func (p *Processor) structToSchema(pkg *packages.Package, st *types.Struct, name string) *model.Schema {
	schema := &model.Schema{
		Type:       "object",
		Title:      name,
		Properties: make(map[string]*model.Schema),
	}

	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		tag := st.Tag(i)

		// 跳过未导出字段
		if !field.Exported() {
			continue
		}

		// 解析 json tag
		jsonName, omitempty := parseJSONTag(tag)
		if jsonName == "-" {
			continue
		}
		if jsonName == "" {
			jsonName = field.Name()
		}

		// 解析字段类型
		fieldSchema := p.typeToSchema(field.Type())

		// 添加字段描述 (从 tag 或 注释)
		if desc := parseDescTag(tag); desc != "" {
			fieldSchema.Description = desc
		}

		schema.Properties[jsonName] = fieldSchema

		// 处理必填字段
		if !omitempty {
			schema.Required = append(schema.Required, jsonName)
		}
	}

	return schema
}

// typeToSchema 将 Go 类型转换为 Schema
func (p *Processor) typeToSchema(t types.Type) *model.Schema {
	switch typ := t.(type) {
	case *types.Basic:
		return p.basicTypeToSchema(typ)

	case *types.Pointer:
		return p.typeToSchema(typ.Elem())

	case *types.Slice:
		return &model.Schema{
			Type:  "array",
			Items: p.typeToSchema(typ.Elem()),
		}

	case *types.Array:
		return &model.Schema{
			Type:     "array",
			Items:    p.typeToSchema(typ.Elem()),
			MaxItems: int(typ.Len()),
			MinItems: int(typ.Len()),
		}

	case *types.Map:
		return &model.Schema{
			Type:                 "object",
			AdditionalProperties: p.typeToSchema(typ.Elem()),
		}

	case *types.Named:
		// 递归解析命名类型
		obj := typ.Obj()
		pkgPath := ""
		if obj.Pkg() != nil {
			pkgPath = obj.Pkg().Path()
		}
		fullName := pkgPath + "." + obj.Name()

		// 检查缓存
		if ref, ok := p.GeneratedSchemas[fullName]; ok {
			return &model.Schema{Ref: ref}
		}

		// 处理泛型类型 (Go 1.18+)
		if typeParams := typ.TypeParams(); typeParams != nil && typeParams.Len() > 0 {
			return p.handleGenericNamedType(typ)
		}

		// 解析底层类型
		return p.typeToSchema(typ.Underlying())

	case *types.Struct:
		// 内联匿名结构体
		return p.structToSchema(nil, typ, "")

	case *types.Interface:
		// interface{} -> any
		return &model.Schema{Type: "object"}

	case *types.TypeParam:
		// 泛型类型参数 (T, K, V 等)
		return &model.Schema{
			Ref: "#/components/schemas/" + typ.Obj().Name(),
		}

	default:
		return &model.Schema{Type: "object"}
	}
}

// handleGenericNamedType 处理泛型命名类型
func (p *Processor) handleGenericNamedType(typ *types.Named) *model.Schema {
	// 获取类型参数
	typeArgs := typ.TypeArgs()
	if typeArgs == nil || typeArgs.Len() == 0 {
		return p.typeToSchema(typ.Underlying())
	}

	// 构建类型参数字符串
	var args []string
	for i := 0; i < typeArgs.Len(); i++ {
		arg := typeArgs.At(i)
		if named, ok := arg.(*types.Named); ok {
			args = append(args, named.Obj().Name())
		} else {
			args = append(args, arg.String())
		}
	}

	// 生成 Schema 名称
	baseName := typ.Obj().Name()
	schemaName := generateGenericSchemaName(baseName, args)

	// 检查缓存
	if ref, ok := p.GeneratedSchemas[schemaName]; ok {
		return &model.Schema{Ref: ref}
	}

	// 解析并实例化
	refPath := "#/components/schemas/" + schemaName
	p.GeneratedSchemas[schemaName] = refPath

	// 解析底层结构体
	underlying, ok := typ.Underlying().(*types.Struct)
	if !ok {
		return &model.Schema{Ref: refPath}
	}

	// 创建实例化的 Schema
	schema := &model.Schema{
		Type:       "object",
		Title:      schemaName,
		Properties: make(map[string]*model.Schema),
	}

	for i := 0; i < underlying.NumFields(); i++ {
		field := underlying.Field(i)
		tag := underlying.Tag(i)

		if !field.Exported() {
			continue
		}

		jsonName, omitempty := parseJSONTag(tag)
		if jsonName == "-" {
			continue
		}
		if jsonName == "" {
			jsonName = field.Name()
		}

		// 解析字段类型，替换类型参数
		fieldSchema := p.typeToSchema(field.Type())
		schema.Properties[jsonName] = fieldSchema

		if !omitempty {
			schema.Required = append(schema.Required, jsonName)
		}
	}

	// 添加到 Components
	p.ensureComponents()
	if p.OpenAPI.Components.Schemas == nil {
		p.OpenAPI.Components.Schemas = make(map[string]*model.Schema)
	}
	p.OpenAPI.Components.Schemas[schemaName] = schema

	return &model.Schema{Ref: refPath}
}

// basicTypeToSchema 将 Go 基本类型转换为 Schema
func (p *Processor) basicTypeToSchema(typ *types.Basic) *model.Schema {
	switch typ.Kind() {
	case types.Bool:
		return &model.Schema{Type: "boolean"}
	case types.Int, types.Int8, types.Int16, types.Int32,
		types.Uint, types.Uint8, types.Uint16, types.Uint32:
		return &model.Schema{Type: "integer", Format: "int32"}
	case types.Int64, types.Uint64:
		return &model.Schema{Type: "integer", Format: "int64"}
	case types.Float32:
		return &model.Schema{Type: "number", Format: "float"}
	case types.Float64:
		return &model.Schema{Type: "number", Format: "double"}
	case types.String:
		return &model.Schema{Type: "string"}
	default:
		return &model.Schema{Type: "string"}
	}
}

// primitiveTypeToSchema 将类型名字符串转换为 Schema
func (p *Processor) primitiveTypeToSchema(typeName string) *model.Schema {
	switch strings.ToLower(typeName) {
	case "string":
		return &model.Schema{Type: "string"}
	case "int", "integer", "int32", "uint", "uint32":
		return &model.Schema{Type: "integer", Format: "int32"}
	case "int64", "uint64":
		return &model.Schema{Type: "integer", Format: "int64"}
	case "float", "float32":
		return &model.Schema{Type: "number", Format: "float"}
	case "float64", "double", "number":
		return &model.Schema{Type: "number", Format: "double"}
	case "bool", "boolean":
		return &model.Schema{Type: "boolean"}
	case "file":
		return &model.Schema{Type: "string", Format: "binary"}
	case "time", "time.time":
		return &model.Schema{Type: "string", Format: "date-time"}
	case "date":
		return &model.Schema{Type: "string", Format: "date"}
	case "uuid":
		return &model.Schema{Type: "string", Format: "uuid"}
	case "uri", "url":
		return &model.Schema{Type: "string", Format: "uri"}
	case "email":
		return &model.Schema{Type: "string", Format: "email"}
	case "byte", "bytes":
		return &model.Schema{Type: "string", Format: "byte"}
	case "binary":
		return &model.Schema{Type: "string", Format: "binary"}
	case "any", "interface{}", "object":
		return &model.Schema{Type: "object"}
	default:
		return &model.Schema{}
	}
}

// parseJSONTag 解析 json tag
// 输入: `json:"name,omitempty"`
// 输出: ("name", true)
func parseJSONTag(tag string) (name string, omitempty bool) {
	// 查找 json tag
	jsonTag := ""
	for _, t := range strings.Split(tag, " ") {
		t = strings.Trim(t, "`")
		if strings.HasPrefix(t, "json:") {
			jsonTag = strings.TrimPrefix(t, "json:")
			jsonTag = strings.Trim(jsonTag, "\"")
			break
		}
	}

	if jsonTag == "" {
		return "", false
	}

	parts := strings.Split(jsonTag, ",")
	name = parts[0]

	for _, opt := range parts[1:] {
		if opt == "omitempty" {
			omitempty = true
		}
	}

	return name, omitempty
}

// parseDescTag 解析描述 tag
// 支持: desc:"xxx" 或 description:"xxx"
func parseDescTag(tag string) string {
	for _, prefix := range []string{"desc:", "description:"} {
		if idx := strings.Index(tag, prefix); idx != -1 {
			start := idx + len(prefix)
			if start < len(tag) && tag[start] == '"' {
				end := strings.Index(tag[start+1:], "\"")
				if end != -1 {
					return tag[start+1 : start+1+end]
				}
			}
		}
	}
	return ""
}

// Helper: 不使用的变量避免编译错误
var _ = fmt.Sprint
