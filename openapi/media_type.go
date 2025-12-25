package openapi

// MediaType 提供媒体类型的详细信息
type MediaType struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 定义完整内容的 Schema
	Schema *Schema `yaml:"schema,omitempty"`
	// 定义序列媒体类型中每个项的 Schema
	ItemSchema *Schema `yaml:"itemSchema,omitempty"`
	// 示例值
	Example any `yaml:"example,omitempty"`
	// 多个示例值
	Examples map[string]*Example `yaml:"examples,omitempty"`
	// 按名称编码的映射 (用于 multipart 和 application/x-www-form-urlencoded)
	Encoding map[string]*Encoding `yaml:"encoding,omitempty"`
	// 按位置编码的数组 (用于 multipart)
	PrefixEncoding []*Encoding `yaml:"prefixEncoding,omitempty"`
	// 数组项的编码 (用于 multipart)
	ItemEncoding *Encoding `yaml:"itemEncoding,omitempty"`
}
