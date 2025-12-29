package parser

import (
	"go/ast"
	"strings"

	"github.com/promonkeyli/goas/pkg/model"
	"golang.org/x/tools/go/packages"
)

type Processor struct {
	// 保存加载的所有包信息
	// Key: 包路径 (PkgPath), e.g., "github.com/myproject/models"
	// Value: 包对象
	PackagesMap map[string]*packages.Package

	// 你的 OpenAPI 结果
	OpenAPI *model.T

	// 缓存已生成的 Schema，防止重复和循环引用
	// Key: 全限定类型名 (e.g., "github.com/myproject/models.User")
	GeneratedSchemas map[string]string

	// 当前解析的请求/响应 MIME 类型
	acceptTypes  []string
	produceTypes []string

	// 标记是否已解析全局注释
	globalParsed bool
}

func newProcessor() *Processor {
	return &Processor{
		PackagesMap:      make(map[string]*packages.Package),
		GeneratedSchemas: make(map[string]string),
		OpenAPI:          &model.T{},
	}
}

// addToIndex 递归将包及其依赖加入索引
func (p *Processor) addToIndex(pkg *packages.Package) {
	// 防止重复处理
	if _, exists := p.PackagesMap[pkg.PkgPath]; exists {
		return
	}

	p.PackagesMap[pkg.PkgPath] = pkg

	// 也可以选择递归把 imports 里的包也加进来，
	// 这样即使 dirs 没写 model 包，只要 controller 引用了，我们也能查到。
	for _, importedPkg := range pkg.Imports {
		p.addToIndex(importedPkg)
	}
}

// scanPackage 扫描单个包里的 AST
func (p *Processor) scanPackage(pkg *packages.Package) {
	// 跳过标准库 (可选优化，避免扫描 fmt, net/http 等)
	if isStandardLibrary(pkg.PkgPath) {
		return
	}

	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok || fn.Doc == nil {
				return true
			}

			// 检测函数类型并调用对应的解析逻辑
			if p.isMainFunc(pkg, fn) {
				// main 包的 main 函数 -> 解析全局注释
				if !p.globalParsed {
					p.parseGlobalAnnotations(pkg, file, fn)
					p.globalParsed = true
				}
			} else if p.hasRouterAnnotation(fn) {
				// 有 @Router 注解的函数 -> 解析接口注释
				p.resetMimeTypes()
				p.parseOperation(pkg, file, fn)
			}

			return true
		})
	}
}

// isMainFunc 检查是否是 main 包的 main 函数
func (p *Processor) isMainFunc(pkg *packages.Package, fn *ast.FuncDecl) bool {
	return pkg.Name == "main" && fn.Name.Name == "main"
}

// hasRouterAnnotation 检查函数是否有 @Router 注解
func (p *Processor) hasRouterAnnotation(fn *ast.FuncDecl) bool {
	if fn.Doc == nil {
		return false
	}

	for _, comment := range fn.Doc.List {
		text := strings.ToLower(comment.Text)
		if strings.Contains(text, "@router") {
			return true
		}
	}
	return false
}

// resetMimeTypes 重置 MIME 类型 (每个接口单独处理)
func (p *Processor) resetMimeTypes() {
	p.acceptTypes = nil
	p.produceTypes = nil
}

// isStandardLibrary 检查是否是标准库包
func isStandardLibrary(pkgPath string) bool {
	// 标准库包路径不包含 "."
	// 如 "fmt", "net/http" vs "github.com/xxx"
	if pkgPath == "" {
		return true
	}

	// 检查常见的第三方包前缀
	thirdPartyPrefixes := []string{
		"github.com/",
		"gitlab.com/",
		"bitbucket.org/",
		"gopkg.in/",
		"golang.org/x/",
	}

	for _, prefix := range thirdPartyPrefixes {
		if strings.HasPrefix(pkgPath, prefix) {
			return false
		}
	}

	// 如果路径包含 "."，可能是第三方包
	if strings.Contains(pkgPath, ".") {
		return false
	}

	// 其他情况认为是标准库
	return true
}
