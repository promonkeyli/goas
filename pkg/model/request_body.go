package model

// RequestBody 描述单个请求体
type RequestBody struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 简短描述
	Description string `yaml:"description,omitempty"`
	// 内容及其媒体类型映射
	Content map[string]*MediaType `yaml:"content"`
	// 是否需要请求体
	Required bool `yaml:"required,omitempty"`
}
