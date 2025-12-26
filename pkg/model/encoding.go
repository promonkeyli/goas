package openapi

// Encoding 描述属性如何序列化
type Encoding struct {
	// 内容类型
	ContentType string `yaml:"contentType,omitempty"`
	// 头部映射
	Headers map[string]*Header `yaml:"headers,omitempty"`
	// 序列化风格
	Style string `yaml:"style,omitempty"`
	// 是否展开
	Explode bool `yaml:"explode,omitempty"`
	// 是否允许保留字符
	AllowReserved bool `yaml:"allowReserved,omitempty"`
	// 嵌套编码映射
	Encoding map[string]*Encoding `yaml:"encoding,omitempty"`
	// 嵌套位置编码数组
	PrefixEncoding []*Encoding `yaml:"prefixEncoding,omitempty"`
	// 嵌套项编码
	ItemEncoding *Encoding `yaml:"itemEncoding,omitempty"`
}
