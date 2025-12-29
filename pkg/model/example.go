package model

// Example 描述一个示例对象
type Example struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 简短摘要
	Summary string `yaml:"summary,omitempty"`
	// 长描述
	Description string `yaml:"description,omitempty"`
	// 数据值 (符合 Schema 的数据结构示例)
	DataValue any `yaml:"dataValue,omitempty"`
	// 序列化值 (包含编码和转义的序列化形式)
	SerializedValue string `yaml:"serializedValue,omitempty"`
	// 外部示例的 URI
	ExternalValue string `yaml:"externalValue,omitempty"`
	// 嵌入的示例值 (已弃用，建议使用 dataValue 或 serializedValue)
	Value any `yaml:"value,omitempty"`
}
