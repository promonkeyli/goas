package model

import "encoding/json"

// Paths 持有各个路径及其操作的定义
type Paths struct {
	// 接口的路径映射
	Paths map[string]*PathItem `json:"-"` // JSON 无 inline tag，用 MarshalJSON/UnmarshalJSON 扁平化输出
}

func (p Paths) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Paths)
}

func (p *Paths) UnmarshalJSON(data []byte) error {
	if p == nil {
		return nil
	}
	return json.Unmarshal(data, &p.Paths)
}

// PathItem 描述在单个路径上可用的操作
type PathItem struct {
	// 允许引用此路径项的定义
	Ref string `json:"$ref,omitempty"`
	// 适用于此路径中所有操作的可选摘要
	Summary string `json:"summary,omitempty"`
	// 适用于此路径中所有操作的可选描述
	Description string `json:"description,omitempty"`
	// 此路径上的各种 HTTP 操作定义
	Get     *Operation `json:"get,omitempty"`
	Put     *Operation `json:"put,omitempty"`
	Post    *Operation `json:"post,omitempty"`
	Delete  *Operation `json:"delete,omitempty"`
	Options *Operation `json:"options,omitempty"`
	Head    *Operation `json:"head,omitempty"`
	Patch   *Operation `json:"patch,omitempty"`
	Trace   *Operation `json:"trace,omitempty"`
	Query   *Operation `json:"query,omitempty"`
	// 其他额外操作
	AdditionalOperations map[string]*Operation `json:"additionalOperations,omitempty"`
	// 覆盖全局的服务列表
	Servers []*Server `json:"servers,omitempty"`
	// 适用于此路径下所有操作的参数
	Parameters []*Parameter `json:"parameters,omitempty"`
}

// Operation 描述路径上的单个 API 操作
type Operation struct {
	// 标签列表
	Tags []string `json:"tags,omitempty"`
	// 短摘要
	Summary string `json:"summary,omitempty"`
	// 详细说明
	Description string `json:"description,omitempty"`
	// 外部文档
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
	// 唯一标识符
	OperationID string `json:"operationId,omitempty"`
	// 参数列表
	Parameters []*Parameter `json:"parameters,omitempty"`
	// 请求体
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	// 可能的响应列表
	Responses *Responses `json:"responses"`
	// 回调映射
	Callbacks map[string]*Callback `json:"callbacks,omitempty"`
	// 声明已弃用
	Deprecated bool `json:"deprecated,omitempty"`
	// 安全机制要求
	Security []SecurityRequirement `json:"security,omitempty"`
	// 覆盖全局的服务列表
	Servers []*Server `json:"servers,omitempty"`
}

// Callback 描述一组可能由 API 提供者发起的请求
type Callback struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 键为表达式，值为描述请求的路径项
	Paths map[string]*PathItem `json:"-"`
}

func (c Callback) MarshalJSON() ([]byte, error) {
	out := map[string]any{}
	if c.Ref != "" {
		out["$ref"] = c.Ref
	}
	for k, v := range c.Paths {
		out[k] = v
	}
	return json.Marshal(out)
}

func (c *Callback) UnmarshalJSON(data []byte) error {
	if c == nil {
		return nil
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var paths map[string]*PathItem
	for k, v := range raw {
		if k == "$ref" {
			_ = json.Unmarshal(v, &c.Ref)
			continue
		}

		var pi PathItem
		if err := json.Unmarshal(v, &pi); err != nil {
			return err
		}

		if paths == nil {
			paths = make(map[string]*PathItem, len(raw))
		}
		paths[k] = &pi
	}

	if len(paths) > 0 {
		c.Paths = paths
	} else {
		c.Paths = nil
	}

	return nil
}
