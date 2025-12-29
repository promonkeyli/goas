package parser

import (
	"regexp"
	"strings"
)

// ========== 全局注解标记 ==========

const (
	// 根配置
	TagOpenAPI           = "@openapi"
	TagSelf              = "@self"
	TagJsonSchemaDialect = "@jsonschemadialect"

	// 基本信息
	TagTitle          = "@title"
	TagVersion        = "@version"
	TagSummary        = "@summary"
	TagDescription    = "@description"
	TagTermsOfService = "@termsofservice"

	// 联系人
	TagContactName  = "@contact.name"
	TagContactURL   = "@contact.url"
	TagContactEmail = "@contact.email"

	// 许可证
	TagLicenseName       = "@license.name"
	TagLicenseIdentifier = "@license.identifier"
	TagLicenseURL        = "@license.url"

	// 服务器
	TagServer = "@server"

	// 外部文档
	TagExternalDocs = "@externaldocs"

	// 标签
	TagTagName     = "@tag.name"
	TagTagSummary  = "@tag.summary"
	TagTagDesc     = "@tag.desc"
	TagTagParent   = "@tag.parent"
	TagTagKind     = "@tag.kind"
	TagTagDocsURL  = "@tag.docs.url"
	TagTagDocsDesc = "@tag.docs.desc"

	// 安全方案
	TagSecurityScheme = "@securityscheme"
	TagSecurityScope  = "@securityscope"
	TagSecurity       = "@security"
)

// ========== 接口注解标记 ==========

const (
	// 路由配置
	TagRouter     = "@router"
	TagId         = "@id"
	TagIgnore     = "@ignore"
	TagDeprecated = "@deprecated"

	// 请求控制
	TagParam  = "@param"
	TagAccept = "@accept"

	// 响应定义
	TagSuccess = "@success"
	TagFailure = "@failure"
	TagProduce = "@produce"
	TagHeader  = "@header"

	// 标签
	TagTags = "@tags"
)

// ========== MIME 类型映射 ==========

var mimeTypeMap = map[string]string{
	"json":             "application/json",
	"xml":              "application/xml",
	"plain":            "text/plain",
	"html":             "text/html",
	"form":             "application/x-www-form-urlencoded",
	"multipart":        "multipart/form-data",
	"stream":           "application/octet-stream",
	"application/json": "application/json",
	"application/xml":  "application/xml",
	"text/plain":       "text/plain",
	"text/html":        "text/html",
}

// ========== 工具函数 ==========

// parseCommentLine 解析单行注释，提取注解标记和内容
// 输入: "// @Title My API Title"
// 输出: ("@title", "My API Title")
func parseCommentLine(line string) (tag string, content string) {
	// 去除前缀 "//" 和空格
	line = strings.TrimPrefix(line, "//")
	line = strings.TrimSpace(line)

	if line == "" || !strings.HasPrefix(line, "@") {
		return "", ""
	}

	// 找到第一个空格分隔标记和内容
	parts := strings.SplitN(line, " ", 2)
	tag = strings.ToLower(parts[0])

	if len(parts) > 1 {
		content = strings.TrimSpace(parts[1])
	}

	return tag, content
}

// splitParams 按空格分割参数，但保留引号内的内容
// 输入: `id path int true "用户ID"`
// 输出: ["id", "path", "int", "true", "用户ID"]
func splitParams(s string) []string {
	var result []string
	var current strings.Builder
	inQuote := false
	quoteChar := rune(0)

	for _, r := range s {
		switch {
		case (r == '"' || r == '\'') && !inQuote:
			inQuote = true
			quoteChar = r
		case r == quoteChar && inQuote:
			inQuote = false
			quoteChar = 0
		case r == ' ' && !inQuote:
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

// parseMimeTypes 解析逗号分隔的 MIME 类型
// 输入: "json,xml"
// 输出: ["application/json", "application/xml"]
func parseMimeTypes(s string) []string {
	var result []string
	parts := strings.Split(s, ",")

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if mime, ok := mimeTypeMap[strings.ToLower(p)]; ok {
			result = append(result, mime)
		} else {
			result = append(result, p)
		}
	}

	return result
}

// parseRouterPath 解析路由注解
// 输入: "/users/{id} [get]"
// 输出: ("/users/{id}", "get")
func parseRouterPath(s string) (path string, method string) {
	// 使用正则匹配 [method]
	re := regexp.MustCompile(`\[(\w+)\]`)
	matches := re.FindStringSubmatch(s)

	if len(matches) > 1 {
		method = strings.ToLower(matches[1])
		path = strings.TrimSpace(re.ReplaceAllString(s, ""))
	} else {
		path = strings.TrimSpace(s)
		method = "get" // 默认 GET
	}

	return path, method
}

// parseResponseType 解析响应类型注解
// 输入: "200 {object} model.User \"成功\""
// 输出: (200, "object", "model.User", "成功")
func parseResponseType(s string) (status string, dataType string, dataModel string, desc string) {
	// 使用正则匹配 {type}
	re := regexp.MustCompile(`\{(\w+)\}`)

	params := splitParams(s)
	if len(params) < 1 {
		return
	}

	status = params[0]

	for i := 1; i < len(params); i++ {
		p := params[i]
		if matches := re.FindStringSubmatch(p); len(matches) > 1 {
			dataType = matches[1]
		} else if dataModel == "" && dataType != "" {
			dataModel = p
		} else if dataModel != "" {
			if desc != "" {
				desc += " "
			}
			desc += p
		}
	}

	return status, dataType, dataModel, desc
}

// parseServerLine 解析服务器注解
// 输入: "http://localhost:8080 name=dev 开发环境"
// 输出: (url, name, description)
func parseServerLine(s string) (url, name, description string) {
	parts := splitParams(s)
	if len(parts) < 1 {
		return
	}

	url = parts[0]

	var descParts []string
	for i := 1; i < len(parts); i++ {
		p := parts[i]
		if strings.HasPrefix(p, "name=") {
			name = strings.TrimPrefix(p, "name=")
		} else {
			descParts = append(descParts, p)
		}
	}

	description = strings.Join(descParts, " ")
	return
}
