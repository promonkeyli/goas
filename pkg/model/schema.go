package model

// Schema 允许定义输入和输出数据类型
type Schema struct {
	// JSON Schema 核心字段
	Ref string `yaml:"$ref,omitempty"`
	ID  string `yaml:"$id,omitempty"`

	// 元数据
	Title       string `yaml:"title,omitempty"`
	Description string `yaml:"description,omitempty"`
	Comment     string `yaml:"$comment,omitempty"`

	// 类型定义
	Type   any    `yaml:"type,omitempty"` // 可以是 string 或 []string (3.1+)
	Format string `yaml:"format,omitempty"`
	Const  any    `yaml:"const,omitempty"`
	Enum   []any  `yaml:"enum,omitempty"`

	// 数组相关
	Items       *Schema   `yaml:"items,omitempty"`
	PrefixItems []*Schema `yaml:"prefixItems,omitempty"`
	MaxItems    int       `yaml:"maxItems,omitempty"`
	MinItems    int       `yaml:"minItems,omitempty"`
	UniqueItems bool      `yaml:"uniqueItems,omitempty"`
	Contains    *Schema   `yaml:"contains,omitempty"`

	// 对象相关
	Properties           map[string]*Schema `yaml:"properties,omitempty"`
	PatternProperties    map[string]*Schema `yaml:"patternProperties,omitempty"`
	AdditionalProperties any                `yaml:"additionalProperties,omitempty"` // bool 或 *Schema
	Required             []string           `yaml:"required,omitempty"`
	MaxProperties        int                `yaml:"maxProperties,omitempty"`
	MinProperties        int                `yaml:"minProperties,omitempty"`
	PropertyNames        *Schema            `yaml:"propertyNames,omitempty"`

	// 组合模式
	AllOf []*Schema `yaml:"allOf,omitempty"`
	OneOf []*Schema `yaml:"oneOf,omitempty"`
	AnyOf []*Schema `yaml:"anyOf,omitempty"`
	Not   *Schema   `yaml:"not,omitempty"`

	// 数值限制
	MultipleOf       float64 `yaml:"multipleOf,omitempty"`
	Maximum          any     `yaml:"maximum,omitempty"` // number 或 bool (exclusiveMaximum)
	ExclusiveMaximum any     `yaml:"exclusiveMaximum,omitempty"`
	Minimum          any     `yaml:"minimum,omitempty"` // number 或 bool (exclusiveMinimum)
	ExclusiveMinimum any     `yaml:"exclusiveMinimum,omitempty"`

	// 字符串限制
	MaxLength int    `yaml:"maxLength,omitempty"`
	MinLength int    `yaml:"minLength,omitempty"`
	Pattern   string `yaml:"pattern,omitempty"`

	// 字符串内容
	ContentMediaType string  `yaml:"contentMediaType,omitempty"`
	ContentEncoding  string  `yaml:"contentEncoding,omitempty"`
	ContentSchema    *Schema `yaml:"contentSchema,omitempty"`

	// 默认值和示例
	Default  any   `yaml:"default,omitempty"`
	Examples []any `yaml:"examples,omitempty"`

	// OpenAPI 特定扩展
	Nullable      bool           `yaml:"nullable,omitempty"`
	Discriminator *Discriminator `yaml:"discriminator,omitempty"`
	ReadOnly      bool           `yaml:"readOnly,omitempty"`
	WriteOnly     bool           `yaml:"writeOnly,omitempty"`
	XML           *XML           `yaml:"xml,omitempty"`
	ExternalDocs  *ExternalDocs  `yaml:"externalDocs,omitempty"`
	Example       any            `yaml:"example,omitempty"` // 已弃用，使用 examples
	Deprecated    bool           `yaml:"deprecated,omitempty"`
}

// Discriminator 帮助多态
type Discriminator struct {
	// 要区分的属性名称
	PropertyName string `yaml:"propertyName"`
	// 区分属性值到 Schema 名称或引用的可选映射
	Mapping map[string]string `yaml:"mapping,omitempty"`
	// 当区分属性不存在或无映射时使用的默认 Schema 名称或引用
	DefaultMapping string `yaml:"defaultMapping,omitempty"`
}

// XML 为描述 XML 数据提供额外信息
type XML struct {
	// 节点类型 (element, attribute, text, cdata, none)
	NodeType string `yaml:"nodeType,omitempty"`
	// 元素或属性的名称
	Name string `yaml:"name,omitempty"`
	// 命名空间的 IRI
	Namespace string `yaml:"namespace,omitempty"`
	// 命名空间的前缀
	Prefix string `yaml:"prefix,omitempty"`
	// 声明此属性是否为 XML 属性 (已弃用，使用 nodeType: "attribute")
	Attribute bool `yaml:"attribute,omitempty"`
	// 声明数组模式是否应该包裹在外面 (已弃用，使用 nodeType: "element")
	Wrapped bool `yaml:"wrapped,omitempty"`
}
