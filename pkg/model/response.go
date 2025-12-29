package model

// Responses 操作预期响应的容器
type Responses struct {
	// 默认响应
	Default *Response `yaml:"default,omitempty"`
	// HTTP 状态码与响应对象的映射
	Codes map[string]*Response `yaml:",inline"`
}

// Response 描述 API 操作的单个响应
type Response struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 响应的简短摘要
	Summary string `yaml:"summary,omitempty"`
	// 响应的详细描述
	Description string `yaml:"description"`
	// 响应头映射
	Headers map[string]*Header `yaml:"headers,omitempty"`
	// 响应内容映射
	Content map[string]*MediaType `yaml:"content,omitempty"`
	// 相关的链接映射
	Links map[string]*Link `yaml:"links,omitempty"`
}

// Link 表示一个操作到另一个操作的链接
type Link struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 指向操作的引用或 ID
	OperationRef string `yaml:"operationRef,omitempty"`
	OperationID  string `yaml:"operationId,omitempty"`
	// 传递的参数和请求体
	Parameters  map[string]any `yaml:"parameters,omitempty"`
	RequestBody any            `yaml:"requestBody,omitempty"`
	// 说明与服务器覆盖
	Description string  `yaml:"description,omitempty"`
	Server      *Server `yaml:"server,omitempty"`
}
