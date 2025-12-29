package model

// Example 描述一个示例对象
type Example struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 简短摘要
	Summary string `json:"summary,omitempty"`
	// 长描述
	Description string `json:"description,omitempty"`
	// 数据值 (符合 Schema 的数据结构示例)
	DataValue any `json:"dataValue,omitempty"`
	// 序列化值 (包含编码和转义的序列化形式)
	SerializedValue string `json:"serializedValue,omitempty"`
	// 外部示例的 URI
	ExternalValue string `json:"externalValue,omitempty"`
	// 嵌入的示例值 (已弃用，建议使用 dataValue 或 serializedValue)
	Value any `json:"value,omitempty"`
}
