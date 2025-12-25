package openapi

type Path struct {
	// 引用
	Ref string `yaml:"$ref,omitempty"`
	// 概述
	Summary string `yaml:"summary,omitempty"`
	// 详情
	Description string `yaml:"description,omitempty"`
	// GET 方法
	Get *Operation `yaml:"get,omitempty"`
	// PUT 方法
	Put *Operation `yaml:"put,omitempty"`
	// POST 方法
	Post *Operation `yaml:"post,omitempty"`
	// DELETE 方法
	Delete *Operation `yaml:"delete,omitempty"`
	// OPTIONS 方法
	Options *Operation `yaml:"options,omitempty"`
	// HEAD 方法
	Head *Operation `yaml:"head,omitempty"`
	// PATCH 方法
	Patch *Operation `yaml:"patch,omitempty"`
	// TRACE 方法
	Trace *Operation `yaml:"trace,omitempty"`
	// query 参数
	Query *Operation `yaml:"query,omitempty"`
	// 其他操作
	AdditionalOperations map[string]*Operation `yaml:",additionalOperations,omitempty"`
	// 服务列表
	Servers []*Server `yaml:"servers,omitempty"`
	// 参数列表
	Parameters []*Parameter `yaml:"parameters,omitempty"`
}

type Operation struct {
}

type Parameter struct {
}
