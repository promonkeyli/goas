package parser

import (
	"go/ast"
	"strings"

	"github.com/promonkeyli/goas/pkg/model"
	"golang.org/x/tools/go/packages"
)

// parseGlobalAnnotations 解析 main 包中 main 函数的全局注释
func (p *Processor) parseGlobalAnnotations(pkg *packages.Package, file *ast.File, fn *ast.FuncDecl) {
	if fn.Doc == nil {
		return
	}

	// 当前正在解析的 Tag (用于处理 @Tag.* 分组)
	var currentTag *model.Tag

	for _, comment := range fn.Doc.List {
		tag, content := parseCommentLine(comment.Text)
		if tag == "" {
			continue
		}

		switch tag {
		// ========== 根配置 ==========
		case TagOpenAPI:
			p.OpenAPI.OpenAPI = content
		case TagSelf:
			p.OpenAPI.Self = content
		case TagJsonSchemaDialect:
			p.OpenAPI.JSONSchemaDialect = content

		// ========== 基本信息 ==========
		case TagTitle:
			p.OpenAPI.Info.Title = content
		case TagVersion:
			p.OpenAPI.Info.Version = content
		case TagSummary:
			p.OpenAPI.Info.Summary = content
		case TagDescription:
			p.appendDescription(&p.OpenAPI.Info.Description, content)
		case TagTermsOfService:
			p.OpenAPI.Info.TermsOfService = content

		// ========== 联系人 ==========
		case TagContactName:
			p.ensureContact()
			p.OpenAPI.Info.Contact.Name = content
		case TagContactURL:
			p.ensureContact()
			p.OpenAPI.Info.Contact.URL = content
		case TagContactEmail:
			p.ensureContact()
			p.OpenAPI.Info.Contact.Email = content

		// ========== 许可证 ==========
		case TagLicenseName:
			p.ensureLicense()
			p.OpenAPI.Info.License.Name = content
		case TagLicenseIdentifier:
			p.ensureLicense()
			p.OpenAPI.Info.License.Identifier = content
		case TagLicenseURL:
			p.ensureLicense()
			p.OpenAPI.Info.License.URL = content

		// ========== 服务器 ==========
		case TagServer:
			url, name, desc := parseServerLine(content)
			server := &model.Server{
				URL:         url,
				Name:        name,
				Description: desc,
			}
			p.OpenAPI.Servers = append(p.OpenAPI.Servers, server)

		// ========== 外部文档 ==========
		case TagExternalDocs:
			p.parseExternalDocs(content)

		// ========== 标签 ==========
		case TagTagName:
			// 开始一个新的 Tag 组
			currentTag = &model.Tag{Name: content}
			p.OpenAPI.Tags = append(p.OpenAPI.Tags, currentTag)
		case TagTagSummary:
			if currentTag != nil {
				currentTag.Summary = content
			}
		case TagTagDesc:
			if currentTag != nil {
				currentTag.Description = content
			}
		case TagTagParent:
			if currentTag != nil {
				currentTag.Parent = content
			}
		case TagTagKind:
			if currentTag != nil {
				currentTag.Kind = content
			}
		case TagTagDocsURL:
			if currentTag != nil {
				p.ensureTagExternalDocs(currentTag)
				currentTag.ExternalDocs.URL = content
			}
		case TagTagDocsDesc:
			if currentTag != nil {
				p.ensureTagExternalDocs(currentTag)
				currentTag.ExternalDocs.Description = content
			}

		// ========== 安全方案 ==========
		case TagSecurityScheme:
			p.parseSecurityScheme(content)
		case TagSecurityScope:
			p.parseSecurityScope(content)
		case TagSecurity:
			p.parseGlobalSecurity(content)
		}
	}
}

// ========== 辅助方法 ==========

func (p *Processor) ensureContact() {
	if p.OpenAPI.Info.Contact == nil {
		p.OpenAPI.Info.Contact = &model.Contact{}
	}
}

func (p *Processor) ensureLicense() {
	if p.OpenAPI.Info.License == nil {
		p.OpenAPI.Info.License = &model.License{}
	}
}

func (p *Processor) ensureComponents() {
	if p.OpenAPI.Components == nil {
		p.OpenAPI.Components = &model.Components{}
	}
}

func (p *Processor) ensureTagExternalDocs(tag *model.Tag) {
	if tag.ExternalDocs == nil {
		tag.ExternalDocs = &model.ExternalDocs{}
	}
}

func (p *Processor) appendDescription(desc *string, content string) {
	if *desc != "" {
		*desc += "\n"
	}
	*desc += content
}

// parseExternalDocs 解析外部文档注解
// 格式: <url> [description]
func (p *Processor) parseExternalDocs(content string) {
	parts := splitParams(content)
	if len(parts) < 1 {
		return
	}

	p.OpenAPI.ExternalDocs = &model.ExternalDocs{
		URL: parts[0],
	}

	if len(parts) > 1 {
		p.OpenAPI.ExternalDocs.Description = strings.Join(parts[1:], " ")
	}
}

// parseSecurityScheme 解析安全方案定义
// 格式: <name> <type> [args...]
// 示例:
//   - apiKey: Key apiKey header Token
//   - http:   JWT http bearer JWT
//   - oauth2: OAuth oauth2 implicit http://auth
func (p *Processor) parseSecurityScheme(content string) {
	params := splitParams(content)
	if len(params) < 2 {
		return
	}

	name := params[0]
	schemeType := strings.ToLower(params[1])

	scheme := &model.SecurityScheme{
		Type: schemeType,
	}

	switch schemeType {
	case "apikey":
		scheme.Type = "apiKey"
		if len(params) > 2 {
			scheme.In = params[2] // header, query, cookie
		}
		if len(params) > 3 {
			scheme.Name = params[3] // 参数名称
		}
	case "http":
		if len(params) > 2 {
			scheme.Scheme = params[2] // bearer, basic
		}
		if len(params) > 3 {
			scheme.BearerFormat = params[3] // JWT
		}
	case "oauth2":
		if len(params) > 2 {
			flowType := params[2] // implicit, password, clientCredentials, authorizationCode
			if len(params) > 3 {
				authURL := params[3]
				scheme.Flows = p.createOAuthFlow(flowType, authURL)
			}
		}
	case "openidconnect":
		scheme.Type = "openIdConnect"
		if len(params) > 2 {
			scheme.OpenIDConnectURL = params[2]
		}
	case "mutualtls":
		scheme.Type = "mutualTLS"
	}

	p.ensureComponents()
	if p.OpenAPI.Components.SecuritySchemes == nil {
		p.OpenAPI.Components.SecuritySchemes = make(map[string]*model.SecurityScheme)
	}
	p.OpenAPI.Components.SecuritySchemes[name] = scheme
}

// createOAuthFlow 创建 OAuth 流程配置
func (p *Processor) createOAuthFlow(flowType, authURL string) *model.OAuthFlows {
	flow := &model.OAuthFlow{
		Scopes: make(map[string]string),
	}

	flows := &model.OAuthFlows{}

	switch flowType {
	case "implicit":
		flow.AuthorizationURL = authURL
		flows.Implicit = flow
	case "password":
		flow.TokenURL = authURL
		flows.Password = flow
	case "clientcredentials", "clientCredentials":
		flow.TokenURL = authURL
		flows.ClientCredentials = flow
	case "authorizationcode", "authorizationCode":
		flow.AuthorizationURL = authURL
		flows.AuthorizationCode = flow
	}

	return flows
}

// parseSecurityScope 解析 OAuth2 Scope
// 格式: <scheme> <scope> <description>
func (p *Processor) parseSecurityScope(content string) {
	params := splitParams(content)
	if len(params) < 3 {
		return
	}

	schemeName := params[0]
	scopeName := params[1]
	scopeDesc := strings.Join(params[2:], " ")

	p.ensureComponents()
	if p.OpenAPI.Components.SecuritySchemes == nil {
		return
	}

	scheme, ok := p.OpenAPI.Components.SecuritySchemes[schemeName]
	if !ok || scheme.Flows == nil {
		return
	}

	// 添加 scope 到所有已配置的流程
	addScopeToFlow := func(flow *model.OAuthFlow) {
		if flow != nil {
			if flow.Scopes == nil {
				flow.Scopes = make(map[string]string)
			}
			flow.Scopes[scopeName] = scopeDesc
		}
	}

	addScopeToFlow(scheme.Flows.Implicit)
	addScopeToFlow(scheme.Flows.Password)
	addScopeToFlow(scheme.Flows.ClientCredentials)
	addScopeToFlow(scheme.Flows.AuthorizationCode)
}

// parseGlobalSecurity 解析全局安全要求
// 格式: <name> [scopes...]
func (p *Processor) parseGlobalSecurity(content string) {
	params := splitParams(content)
	if len(params) < 1 {
		return
	}

	name := params[0]
	scopes := params[1:]

	requirement := model.SecurityRequirement{
		name: scopes,
	}

	p.OpenAPI.Security = append(p.OpenAPI.Security, requirement)
}
