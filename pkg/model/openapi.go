package model

// T 是 OpenAPI 文档的根对象
type T struct {
	// OpenAPI 规范版本号
	OpenAPI string `json:"openapi"`
	// 文档的基础 URI
	Self string `json:"$self,omitempty"`
	// API 的元数据信息
	Info Info `json:"info"`
	// Schema 对象的默认方言
	JSONSchemaDialect string `json:"jsonSchemaDialect,omitempty"`
	// 服务的连接信息列表
	Servers []*Server `json:"servers,omitempty"`
	// API 的可用路径及操作
	Paths *Paths `json:"paths,omitempty"`
	// 接收的 Webhooks 列表
	Webhooks map[string]*PathItem `json:"webhooks,omitempty"`
	// 持有各种可复用的对象
	Components *Components `json:"components,omitempty"`
	// 安全机制声明
	Security []SecurityRequirement `json:"security,omitempty"`
	// 标签列表及元数据
	Tags []*Tag `json:"tags,omitempty"`
	// 外部文档
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}

// Info 提供 API 的元数据
type Info struct {
	// API 标题
	Title string `json:"title"`
	// API 的短摘要
	Summary string `json:"summary,omitempty"`
	// API 的详细描述
	Description string `json:"description,omitempty"`
	// 服务条款的 URI
	TermsOfService string `json:"termsOfService,omitempty"`
	// 联系信息
	Contact *Contact `json:"contact,omitempty"`
	// 许可证信息
	License *License `json:"license,omitempty"`
	// OpenAPI 文档的版本
	Version string `json:"version"`
}

// Contact API 的联系信息
type Contact struct {
	// 联系人或组织名称
	Name string `json:"name,omitempty"`
	// 联系信息的 URI
	URL string `json:"url,omitempty"`
	// 联系人或组织的邮箱地址
	Email string `json:"email,omitempty"`
}

// License API 的许可证信息
type License struct {
	// 许可证名称
	Name string `json:"name"`
	// 许可证的 SPDX 表达式
	Identifier string `json:"identifier,omitempty"`
	// 许可证的 URI
	URL string `json:"url,omitempty"`
}

// Tag 标记使用的标签及元数据
type Tag struct {
	// 标签名称
	Name string `json:"name"`
	// 标签的短摘要
	Summary string `json:"summary,omitempty"`
	// 标签的详细说明
	Description string `json:"description,omitempty"`
	// 额外的外部文档
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
	// 此标签嵌套在其下的父标签名称
	Parent string `json:"parent,omitempty"`
	// 标签类型的机器可读字符串 (如 nav, badge, audience)
	Kind string `json:"kind,omitempty"`
}

// ExternalDocs 引用外部资源以获取扩展文档
type ExternalDocs struct {
	// 目标文档的描述
	Description string `json:"description,omitempty"`
	// 目标文档的 URI
	URL string `json:"url"`
}

// Components 持有各种可复用的对象
type Components struct {
	// 可复用的 Schema 对象
	Schemas map[string]*Schema `json:"schemas,omitempty"`
	// 可复用的响应对象
	Responses map[string]*Response `json:"responses,omitempty"`
	// 可复用的参数对象
	Parameters map[string]*Parameter `json:"parameters,omitempty"`
	// 可复用的示例对象
	Examples map[string]*Example `json:"examples,omitempty"`
	// 可复用的请求体对象
	RequestBodies map[string]*RequestBody `json:"requestBodies,omitempty"`
	// 可复用的头部对象
	Headers map[string]*Header `json:"headers,omitempty"`
	// 可复用的安全方案对象
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
	// 可复用的链接对象
	Links map[string]*Link `json:"links,omitempty"`
	// 可复用的回调对象
	Callbacks map[string]*Callback `json:"callbacks,omitempty"`
	// 可复用的路径项对象
	PathItems map[string]*PathItem `json:"pathItems,omitempty"`
	// 可复用的媒体类型对象
	MediaTypes map[string]*MediaType `json:"mediaTypes,omitempty"`
}

// Reference 允许引用文档内部或外部的其他组件
type Reference struct {
	// 引用标识符，必须是 URI 格式
	Ref string `json:"$ref"`
	// 可选的简短摘要，用于覆盖被引用组件的摘要
	Summary string `json:"summary,omitempty"`
	// 可选的描述，用于覆盖被引用组件的描述
	Description string `json:"description,omitempty"`
}
