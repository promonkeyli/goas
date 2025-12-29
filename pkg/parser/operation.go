package parser

import (
	"go/ast"
	"strings"

	"github.com/promonkeyli/goas/pkg/model"
	"golang.org/x/tools/go/packages"
)

// parseOperation 解析 Handler 函数上的接口注释，生成 Operation 对象
func (p *Processor) parseOperation(pkg *packages.Package, file *ast.File, fn *ast.FuncDecl) {
	if fn.Doc == nil {
		return
	}

	var (
		routerPath   string
		routerMethod string
		ignore       bool
		op           = &model.Operation{
			Responses: &model.Responses{
				Codes: make(map[string]*model.Response),
			},
		}
	)

	for _, comment := range fn.Doc.List {
		tag, content := parseCommentLine(comment.Text)
		if tag == "" {
			continue
		}

		switch tag {
		// ========== 路由配置 ==========
		case TagRouter:
			routerPath, routerMethod = parseRouterPath(content)
		case TagId:
			op.OperationID = content
		case TagIgnore:
			ignore = true
			return
		case TagDeprecated:
			op.Deprecated = true

		// ========== 基本信息 ==========
		case TagSummary:
			op.Summary = content
		case TagDescription:
			if op.Description != "" {
				op.Description += "\n"
			}
			op.Description += content
		case TagTags:
			op.Tags = parseTags(content)

		// ========== 请求控制 ==========
		case TagParam:
			p.parseParam(pkg, op, content)
		case TagAccept:
			// Accept 类型会在 parseParam 处理 body 时使用
			p.acceptTypes = parseMimeTypes(content)

		// ========== 响应定义 ==========
		case TagSuccess, TagFailure:
			p.parseResponse(pkg, op, content)
		case TagProduce:
			p.produceTypes = parseMimeTypes(content)
		case TagHeader:
			p.parseResponseHeader(op, content)

		// ========== 外部文档 ==========
		case TagExternalDocs:
			parts := splitParams(content)
			if len(parts) > 0 {
				op.ExternalDocs = &model.ExternalDocs{URL: parts[0]}
				if len(parts) > 1 {
					op.ExternalDocs.Description = strings.Join(parts[1:], " ")
				}
			}

		// ========== 安全 ==========
		case TagSecurity:
			params := splitParams(content)
			if len(params) > 0 {
				req := model.SecurityRequirement{
					params[0]: params[1:],
				}
				op.Security = append(op.Security, req)
			}
		}
	}

	// 跳过标记为忽略或没有路由的函数
	if ignore || routerPath == "" {
		return
	}

	// 设置默认 OperationID
	if op.OperationID == "" {
		op.OperationID = fn.Name.Name
	}

	// 添加到 Paths
	p.addOperation(routerPath, routerMethod, op)
}

// parseTags 解析标签列表
// 输入: "user, admin"
// 输出: ["user", "admin"]
func parseTags(content string) []string {
	var tags []string
	parts := strings.Split(content, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			tags = append(tags, p)
		}
	}
	return tags
}

// parseParam 解析参数注解
// 格式: <name> <in> <type> <required> <description>
// in: path, query, header, cookie, body, formData
func (p *Processor) parseParam(pkg *packages.Package, op *model.Operation, content string) {
	params := splitParams(content)
	if len(params) < 4 {
		return
	}

	name := params[0]
	in := strings.ToLower(params[1])
	typeName := params[2]
	required := strings.ToLower(params[3]) == "true"

	var desc string
	if len(params) > 4 {
		desc = strings.Join(params[4:], " ")
	}

	// 处理 body 参数
	if in == "body" {
		p.parseBodyParam(pkg, op, typeName, desc, required)
		return
	}

	// 处理 formData 参数
	if in == "formdata" {
		p.parseFormDataParam(pkg, op, name, typeName, desc, required)
		return
	}

	// 处理普通参数 (path, query, header, cookie)
	param := &model.Parameter{
		Name:        name,
		In:          in,
		Description: desc,
		Required:    required || in == "path", // path 参数始终必填
		Schema:      p.primitiveTypeToSchema(typeName),
	}

	op.Parameters = append(op.Parameters, param)
}

// parseBodyParam 解析 body 参数
func (p *Processor) parseBodyParam(pkg *packages.Package, op *model.Operation, typeName, desc string, required bool) {
	schema := p.resolveTypeSchema(pkg, typeName)

	// 确定 content-type
	contentTypes := p.acceptTypes
	if len(contentTypes) == 0 {
		contentTypes = []string{"application/json"}
	}

	content := make(map[string]*model.MediaType)
	for _, ct := range contentTypes {
		content[ct] = &model.MediaType{
			Schema: schema,
		}
	}

	op.RequestBody = &model.RequestBody{
		Description: desc,
		Required:    required,
		Content:     content,
	}
}

// parseFormDataParam 解析 formData 参数
func (p *Processor) parseFormDataParam(pkg *packages.Package, op *model.Operation, name, typeName, desc string, required bool) {
	// 确保 RequestBody 存在
	if op.RequestBody == nil {
		op.RequestBody = &model.RequestBody{
			Content: make(map[string]*model.MediaType),
		}
	}

	// 使用 multipart/form-data 或 application/x-www-form-urlencoded
	contentType := "multipart/form-data"
	if typeName != "file" && len(p.acceptTypes) > 0 {
		for _, ct := range p.acceptTypes {
			if ct == "application/x-www-form-urlencoded" {
				contentType = ct
				break
			}
		}
	}

	mediaType, ok := op.RequestBody.Content[contentType]
	if !ok {
		mediaType = &model.MediaType{
			Schema: &model.Schema{
				Type:       "object",
				Properties: make(map[string]*model.Schema),
			},
		}
		op.RequestBody.Content[contentType] = mediaType
	}

	// 添加属性
	var propSchema *model.Schema
	if typeName == "file" {
		propSchema = &model.Schema{
			Type:   "string",
			Format: "binary",
		}
	} else {
		propSchema = p.primitiveTypeToSchema(typeName)
	}
	propSchema.Description = desc

	mediaType.Schema.Properties[name] = propSchema

	if required {
		mediaType.Schema.Required = append(mediaType.Schema.Required, name)
	}
}

// parseResponse 解析响应注解
// 格式: <status> {<type>} <data> [description]
func (p *Processor) parseResponse(pkg *packages.Package, op *model.Operation, content string) {
	status, dataType, dataModel, desc := parseResponseType(content)
	if status == "" {
		return
	}

	var schema *model.Schema

	switch dataType {
	case "object":
		schema = p.resolveTypeSchema(pkg, dataModel)
	case "array":
		itemSchema := p.resolveTypeSchema(pkg, dataModel)
		schema = &model.Schema{
			Type:  "array",
			Items: itemSchema,
		}
	case "string", "integer", "number", "boolean":
		schema = &model.Schema{Type: dataType}
	default:
		// 如果没有指定类型，尝试解析 dataModel
		if dataModel != "" {
			schema = p.resolveTypeSchema(pkg, dataModel)
		}
	}

	// 确定 content-type
	contentTypes := p.produceTypes
	if len(contentTypes) == 0 {
		contentTypes = []string{"application/json"}
	}

	response := &model.Response{
		Description: desc,
	}

	if schema != nil {
		response.Content = make(map[string]*model.MediaType)
		for _, ct := range contentTypes {
			response.Content[ct] = &model.MediaType{
				Schema: schema,
			}
		}
	}

	// 添加到响应列表
	if status == "default" {
		op.Responses.Default = response
	} else {
		op.Responses.Codes[status] = response
	}
}

// parseResponseHeader 解析响应头注解
// 格式: <status> {<type>} <name> <description>
func (p *Processor) parseResponseHeader(op *model.Operation, content string) {
	params := splitParams(content)
	if len(params) < 3 {
		return
	}

	status := params[0]
	headerType := ""
	headerName := ""
	desc := ""

	for i := 1; i < len(params); i++ {
		p := params[i]
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			headerType = strings.Trim(p, "{}")
		} else if headerName == "" {
			headerName = p
		} else {
			if desc != "" {
				desc += " "
			}
			desc += p
		}
	}

	header := &model.Header{
		Description: desc,
		Schema:      p.primitiveTypeToSchema(headerType),
	}

	// 找到对应的响应并添加 Header
	var resp *model.Response
	if status == "default" {
		resp = op.Responses.Default
	} else {
		resp = op.Responses.Codes[status]
	}

	if resp != nil {
		if resp.Headers == nil {
			resp.Headers = make(map[string]*model.Header)
		}
		resp.Headers[headerName] = header
	}
}

// addOperation 添加操作到 Paths
func (p *Processor) addOperation(path, method string, op *model.Operation) {
	// 确保 Paths 存在
	if p.OpenAPI.Paths == nil {
		p.OpenAPI.Paths = &model.Paths{
			Paths: make(map[string]*model.PathItem),
		}
	}

	// 确保 PathItem 存在
	pathItem, ok := p.OpenAPI.Paths.Paths[path]
	if !ok {
		pathItem = &model.PathItem{}
		p.OpenAPI.Paths.Paths[path] = pathItem
	}

	// 根据方法设置操作
	switch method {
	case "get":
		pathItem.Get = op
	case "post":
		pathItem.Post = op
	case "put":
		pathItem.Put = op
	case "delete":
		pathItem.Delete = op
	case "patch":
		pathItem.Patch = op
	case "head":
		pathItem.Head = op
	case "options":
		pathItem.Options = op
	case "trace":
		pathItem.Trace = op
	case "query":
		pathItem.Query = op
	}
}
