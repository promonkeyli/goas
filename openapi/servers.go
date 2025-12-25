package openapi

type Server struct {
	// 服务名称
	Name string `yaml:"name,omitempty"`
	// 服务地址
	URL string `yaml:"url"`
	// 服务描述
	Description string `yaml:"description,omitempty"`
	// 变量列表：TODO：后续实现这个，用处不大
	// Variables map[string]*ServerVariable `yaml:"variables,omitempty"`
}

type ServerVariable struct {
	// 默认值
	Default string `yaml:"default"`
	// 描述
	Description string `yaml:"description,omitempty"`
	// 可选值列表
	Enum []string `yaml:"enum,omitempty"`
}
