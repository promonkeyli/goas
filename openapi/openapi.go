package openapi

type T struct {
	// OpenAPI 规范版本
	OpenAPI string `yaml:"openapi"`
	// 文档元数据
	Info Info `yaml:"info"`
	// 服务列表
	Servers []*Server `yaml:"servers,omitempty"`
	// 路径列表
	Paths map[string]Path `yaml:"paths"`
}
