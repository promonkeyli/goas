package model

// Schema 允许定义输入和输出数据类型
type Schema struct {
	// JSON Schema 核心字段
	Ref string `json:"$ref,omitempty"`
	ID  string `json:"$id,omitempty"`

	// 元数据
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Comment     string `json:"$comment,omitempty"`

	// 类型定义
	Type   any    `json:"type,omitempty"` // 可以是 string 或 []string (3.1+)
	Format string `json:"format,omitempty"`
	Const  any    `json:"const,omitempty"`
	Enum   []any  `json:"enum,omitempty"`

	// 数组相关
	Items       *Schema   `json:"items,omitempty"`
	PrefixItems []*Schema `json:"prefixItems,omitempty"`
	MaxItems    int       `json:"maxItems,omitempty"`
	MinItems    int       `json:"minItems,omitempty"`
	UniqueItems bool      `json:"uniqueItems,omitempty"`
	Contains    *Schema   `json:"contains,omitempty"`

	// 对象相关
	Properties           map[string]*Schema `json:"properties,omitempty"`
	PatternProperties    map[string]*Schema `json:"patternProperties,omitempty"`
	AdditionalProperties any                `json:"additionalProperties,omitempty"` // bool 或 *Schema
	Required             []string           `json:"required,omitempty"`
	MaxProperties        int                `json:"maxProperties,omitempty"`
	MinProperties        int                `json:"minProperties,omitempty"`
	PropertyNames        *Schema            `json:"propertyNames,omitempty"`

	// 组合模式
	AllOf []*Schema `json:"allOf,omitempty"`
	OneOf []*Schema `json:"oneOf,omitempty"`
	AnyOf []*Schema `json:"anyOf,omitempty"`
	Not   *Schema   `json:"not,omitempty"`

	// 数值限制
	MultipleOf       float64 `json:"multipleOf,omitempty"`
	Maximum          any     `json:"maximum,omitempty"` // number 或 bool (exclusiveMaximum)
	ExclusiveMaximum any     `json:"exclusiveMaximum,omitempty"`
	Minimum          any     `json:"minimum,omitempty"` // number 或 bool (exclusiveMinimum)
	ExclusiveMinimum any     `json:"exclusiveMinimum,omitempty"`

	// 字符串限制
	MaxLength int    `json:"maxLength,omitempty"`
	MinLength int    `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`

	// 字符串内容
	ContentMediaType string  `json:"contentMediaType,omitempty"`
	ContentEncoding  string  `json:"contentEncoding,omitempty"`
	ContentSchema    *Schema `json:"contentSchema,omitempty"`

	// 默认值和示例
	Default  any   `json:"default,omitempty"`
	Examples []any `json:"examples,omitempty"`

	// OpenAPI 特定扩展
	Nullable      bool           `json:"nullable,omitempty"`
	Discriminator *Discriminator `json:"discriminator,omitempty"`
	ReadOnly      bool           `json:"readOnly,omitempty"`
	WriteOnly     bool           `json:"writeOnly,omitempty"`
	XML           *XML           `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocs  `json:"externalDocs,omitempty"`
	Example       any            `json:"example,omitempty"` // 已弃用，使用 examples
	Deprecated    bool           `json:"deprecated,omitempty"`
}

// Discriminator 帮助多态
type Discriminator struct {
	// 要区分的属性名称
	PropertyName string `json:"propertyName"`
	// 区分属性值到 Schema 名称或引用的可选映射
	Mapping map[string]string `json:"mapping,omitempty"`
	// 当区分属性不存在或无映射时使用的默认 Schema 名称或引用
	DefaultMapping string `json:"defaultMapping,omitempty"`
}

// XML 为描述 XML 数据提供额外信息
type XML struct {
	// 节点类型 (element, attribute, text, cdata, none)
	NodeType string `json:"nodeType,omitempty"`
	// 元素或属性的名称
	Name string `json:"name,omitempty"`
	// 命名空间的 IRI
	Namespace string `json:"namespace,omitempty"`
	// 命名空间的前缀
	Prefix string `json:"prefix,omitempty"`
	// 声明此属性是否为 XML 属性 (已弃用，使用 nodeType: "attribute")
	Attribute bool `json:"attribute,omitempty"`
	// 声明数组模式是否应该包裹在外面 (已弃用，使用 nodeType: "element")
	Wrapped bool `json:"wrapped,omitempty"`
}
