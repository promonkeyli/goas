package model

// SecurityScheme 定义一个安全方案
type SecurityScheme struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 类型 (apiKey, http, mutualTLS, oauth2, openIdConnect)
	Type string `yaml:"type"`
	// 简短描述
	Description string `yaml:"description,omitempty"`
	// 参数名称与位置 (仅 apiKey)
	Name string `yaml:"name,omitempty"`
	In   string `yaml:"in,omitempty"`
	// 授权方案与格式 (仅 http)
	Scheme       string `yaml:"scheme,omitempty"`
	BearerFormat string `yaml:"bearerFormat,omitempty"`
	// OAuth 流配置
	Flows *OAuthFlows `yaml:"flows,omitempty"`
	// OpenID 配置 URL
	OpenIDConnectURL string `yaml:"openIdConnectUrl,omitempty"`
}

// OAuthFlows 持有支持的 OAuth 流配置
type OAuthFlows struct {
	Implicit          *OAuthFlow `yaml:"implicit,omitempty"`
	Password          *OAuthFlow `yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `yaml:"authorizationCode,omitempty"`
}

// OAuthFlow 描述单个 OAuth 流程
type OAuthFlow struct {
	AuthorizationURL string            `yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `yaml:"tokenUrl,omitempty"`
	RefreshURL       string            `yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `yaml:"scopes"`
}

// SecurityRequirement 列出执行此操作所需的安全性方案
type SecurityRequirement map[string][]string
