package openapi

// Parameter 描述单个操作参数
type Parameter struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 参数名称
	Name string `yaml:"name,omitempty"`
	// 位置 (query, querystring, header, path, cookie)
	In string `yaml:"in,omitempty"`
	// 简短描述
	Description string `yaml:"description,omitempty"`
	// 是否必选
	Required bool `yaml:"required,omitempty"`
	// 是否已弃用
	Deprecated bool `yaml:"deprecated,omitempty"`
	// 是否允许空值
	AllowEmptyValue bool `yaml:"allowEmptyValue,omitempty"`

	// 序列化相关
	Style         string                `yaml:"style,omitempty"`
	Explode       bool                  `yaml:"explode,omitempty"`
	AllowReserved bool                  `yaml:"allowReserved,omitempty"`
	Schema        *Schema               `yaml:"schema,omitempty"`
	Example       any                   `yaml:"example,omitempty"`
	Examples      map[string]*Example   `yaml:"examples,omitempty"`
	Content       map[string]*MediaType `yaml:"content,omitempty"`
}

// Header 描述单个响应头
type Header struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 简短描述
	Description string `yaml:"description,omitempty"`
	// 是否必选
	Required bool `yaml:"required,omitempty"`
	// 是否已弃用
	Deprecated bool `yaml:"deprecated,omitempty"`

	// 序列化相关
	Style    string                `yaml:"style,omitempty"`
	Explode  bool                  `yaml:"explode,omitempty"`
	Schema   *Schema               `yaml:"schema,omitempty"`
	Example  any                   `yaml:"example,omitempty"`
	Examples map[string]*Example   `yaml:"examples,omitempty"`
	Content  map[string]*MediaType `yaml:"content,omitempty"`
}
