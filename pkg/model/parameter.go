package model

// Parameter 描述单个操作参数
type Parameter struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 参数名称
	Name string `json:"name,omitempty"`
	// 位置 (query, querystring, header, path, cookie)
	In string `json:"in,omitempty"`
	// 简短描述
	Description string `json:"description,omitempty"`
	// 是否必选
	Required bool `json:"required,omitempty"`
	// 是否已弃用
	Deprecated bool `json:"deprecated,omitempty"`
	// 是否允许空值
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`

	// 序列化相关
	Style         string                `json:"style,omitempty"`
	Explode       bool                  `json:"explode,omitempty"`
	AllowReserved bool                  `json:"allowReserved,omitempty"`
	Schema        *Schema               `json:"schema,omitempty"`
	Example       any                   `json:"example,omitempty"`
	Examples      map[string]*Example   `json:"examples,omitempty"`
	Content       map[string]*MediaType `json:"content,omitempty"`
}

// Header 描述单个响应头
type Header struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 简短描述
	Description string `json:"description,omitempty"`
	// 是否必选
	Required bool `json:"required,omitempty"`
	// 是否已弃用
	Deprecated bool `json:"deprecated,omitempty"`

	// 序列化相关
	Style    string                `json:"style,omitempty"`
	Explode  bool                  `json:"explode,omitempty"`
	Schema   *Schema               `json:"schema,omitempty"`
	Example  any                   `json:"example,omitempty"`
	Examples map[string]*Example   `json:"examples,omitempty"`
	Content  map[string]*MediaType `json:"content,omitempty"`
}
