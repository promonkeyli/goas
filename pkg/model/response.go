package model

import "encoding/json"

// Responses 操作预期响应的容器
type Responses struct {
	// 默认响应
	Default *Response `json:"default,omitempty"`
	// HTTP 状态码与响应对象的映射
	Codes map[string]*Response `json:"-"`
}

func (r Responses) MarshalJSON() ([]byte, error) {
	out := map[string]*Response{}
	if r.Default != nil {
		out["default"] = r.Default
	}
	for k, v := range r.Codes {
		out[k] = v
	}
	return json.Marshal(out)
}

func (r *Responses) UnmarshalJSON(data []byte) error {
	if r == nil {
		return nil
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var codes map[string]*Response
	for k, v := range raw {
		if k == "default" {
			var resp Response
			if err := json.Unmarshal(v, &resp); err != nil {
				return err
			}
			r.Default = &resp
			continue
		}

		var resp Response
		if err := json.Unmarshal(v, &resp); err != nil {
			return err
		}
		if codes == nil {
			codes = make(map[string]*Response, len(raw))
		}
		codes[k] = &resp
	}

	if len(codes) > 0 {
		r.Codes = codes
	} else {
		r.Codes = nil
	}

	return nil
}

// Response 描述 API 操作的单个响应
type Response struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 响应的简短摘要
	Summary string `json:"summary,omitempty"`
	// 响应的详细描述
	Description string `json:"description"`
	// 响应头映射
	Headers map[string]*Header `json:"headers,omitempty"`
	// 响应内容映射
	Content map[string]*MediaType `json:"content,omitempty"`
	// 相关的链接映射
	Links map[string]*Link `json:"links,omitempty"`
}

// Link 表示一个操作到另一个操作的链接
type Link struct {
	// 允许引用
	Ref string `json:"$ref,omitempty"`
	// 指向操作的引用或 ID
	OperationRef string `json:"operationRef,omitempty"`
	OperationID  string `json:"operationId,omitempty"`
	// 传递的参数和请求体
	Parameters  map[string]any `json:"parameters,omitempty"`
	RequestBody any            `json:"requestBody,omitempty"`
	// 说明与服务器覆盖
	Description string  `json:"description,omitempty"`
	Server      *Server `json:"server,omitempty"`
}
