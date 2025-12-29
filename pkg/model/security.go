package model

// SecurityScheme 定义一个安全方案
type SecurityScheme struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 类型 (apiKey, http, mutualTLS, oauth2, openIdConnect)
	Type string `json:"type"`
	// 简短描述
	Description string `json:"description,omitempty"`
	// 参数名称与位置 (仅 apiKey)
	Name string `json:"name,omitempty"`
	In   string `json:"in,omitempty"`
	// 授权方案与格式 (仅 http)
	Scheme       string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
	// OAuth 流配置
	Flows *OAuthFlows `json:"flows,omitempty"`
	// OpenID 配置 URL
	OpenIDConnectURL string `json:"openIdConnectUrl,omitempty"`
}

// OAuthFlows 持有支持的 OAuth 流配置
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
}

// OAuthFlow 描述单个 OAuth 流程
type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes"`
}

// SecurityRequirement 列出执行此操作所需的安全性方案
type SecurityRequirement map[string][]string
