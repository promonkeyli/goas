package openapi

// Server 表示一个服务器对象
type Server struct {
	// 目标主机的 URL
	URL string `yaml:"url"`
	// 主机的描述
	Description string `yaml:"description,omitempty"`
	// 主机的唯一引用名称
	Name string `yaml:"name,omitempty"`
	// 变量名称及其值的映射，用于 URL 模板替换
	Variables map[string]*ServerVariable `yaml:"variables,omitempty"`
}

// ServerVariable 表示服务器 URL 模板替换的变量
type ServerVariable struct {
	// 用于替换的可选字符串值枚举
	Enum []string `yaml:"enum,omitempty"`
	// 替换时使用的默认值
	Default string `yaml:"default"`
	// 变量的说明描述
	Description string `yaml:"description,omitempty"`
}
