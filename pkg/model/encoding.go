package model

// Encoding 描述属性如何序列化
type Encoding struct {
	// 内容类型
	ContentType string `json:"contentType,omitempty"`
	// 头部映射
	Headers map[string]*Header `json:"headers,omitempty"`
	// 序列化风格
	Style string `json:"style,omitempty"`
	// 是否展开
	Explode bool `json:"explode,omitempty"`
	// 是否允许保留字符
	AllowReserved bool `json:"allowReserved,omitempty"`
	// 嵌套编码映射
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
	// 嵌套位置编码数组
	PrefixEncoding []*Encoding `json:"prefixEncoding,omitempty"`
	// 嵌套项编码
	ItemEncoding *Encoding `json:"itemEncoding,omitempty"`
}
