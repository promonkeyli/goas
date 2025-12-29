package model

// RequestBody 描述单个请求体
type RequestBody struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 简短描述
	Description string `json:"description,omitempty"`
	// 内容及其媒体类型映射
	Content map[string]*MediaType `json:"content"`
	// 是否需要请求体
	Required bool `json:"required,omitempty"`
}
