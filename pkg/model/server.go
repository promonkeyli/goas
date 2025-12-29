package model

// Server 表示一个服务器对象
type Server struct {
	// 目标主机的 URL
	URL string `json:"url"`
	// 主机的描述
	Description string `json:"description,omitempty"`
	// 主机的唯一引用名称
	Name string `json:"name,omitempty"`
	// 变量名称及其值的映射，用于 URL 模板替换
	Variables map[string]*ServerVariable `json:"variables,omitempty"`
}

// ServerVariable 表示服务器 URL 模板替换的变量
type ServerVariable struct {
	// 用于替换的可选字符串值枚举
	Enum []string `json:"enum,omitempty"`
	// 替换时使用的默认值
	Default string `json:"default"`
	// 变量的说明描述
	Description string `json:"description,omitempty"`
}
