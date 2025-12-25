package openapi

// Paths 持有各个路径及其操作的定义
type Paths struct {
	// 接口的路径映射
	Paths map[string]*PathItem `yaml:",inline"`
}

// PathItem 描述在单个路径上可用的操作
type PathItem struct {
	// 允许引用此路径项的定义
	Ref string `yaml:"$ref,omitempty"`
	// 适用于此路径中所有操作的可选摘要
	Summary string `yaml:"summary,omitempty"`
	// 适用于此路径中所有操作的可选描述
	Description string `yaml:"description,omitempty"`
	// 此路径上的各种 HTTP 操作定义
	Get     *Operation `yaml:"get,omitempty"`
	Put     *Operation `yaml:"put,omitempty"`
	Post    *Operation `yaml:"post,omitempty"`
	Delete  *Operation `yaml:"delete,omitempty"`
	Options *Operation `yaml:"options,omitempty"`
	Head    *Operation `yaml:"head,omitempty"`
	Patch   *Operation `yaml:"patch,omitempty"`
	Trace   *Operation `yaml:"trace,omitempty"`
	Query   *Operation `yaml:"query,omitempty"`
	// 其他额外操作
	AdditionalOperations map[string]*Operation `yaml:"additionalOperations,omitempty"`
	// 覆盖全局的服务列表
	Servers []*Server `yaml:"servers,omitempty"`
	// 适用于此路径下所有操作的参数
	Parameters []*Parameter `yaml:"parameters,omitempty"`
}

// Operation 描述路径上的单个 API 操作
type Operation struct {
	// 标签列表
	Tags []string `yaml:"tags,omitempty"`
	// 短摘要
	Summary string `yaml:"summary,omitempty"`
	// 详细说明
	Description string `yaml:"description,omitempty"`
	// 外部文档
	ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty"`
	// 唯一标识符
	OperationID string `yaml:"operationId,omitempty"`
	// 参数列表
	Parameters []*Parameter `yaml:"parameters,omitempty"`
	// 请求体
	RequestBody *RequestBody `yaml:"requestBody,omitempty"`
	// 可能的响应列表
	Responses *Responses `yaml:"responses"`
	// 回调映射
	Callbacks map[string]*Callback `yaml:"callbacks,omitempty"`
	// 声明已弃用
	Deprecated bool `yaml:"deprecated,omitempty"`
	// 安全机制要求
	Security []SecurityRequirement `yaml:"security,omitempty"`
	// 覆盖全局的服务列表
	Servers []*Server `yaml:"servers,omitempty"`
}

// Callback 描述一组可能由 API 提供者发起的请求
type Callback struct {
	// 允许引用
	Ref string `yaml:"$ref,omitempty"`
	// 键为表达式，值为描述请求的路径项
	Paths map[string]*PathItem `yaml:",inline"`
}
