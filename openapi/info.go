package openapi

type Info struct {
	// 标题
	Title string `yaml:"title"`
	// 版本
	Version string `yaml:"version"`
	// 简要概述
	Summary string `yaml:"summary,omitempty"`
	// 详细描述
	Description string `yaml:"description,omitempty"`
	// 服务条款
	TermsOfService string `yaml:"termsOfService,omitempty"`
	// 联系信息
	Contact Contact `yaml:"contact,omitempty"`
	// 许可证信息
	License License `yaml:"license,omitempty"`
}

type Contact struct {
	// 姓名
	Name string `yaml:"name,omitempty"`
	// 网址
	URL string `yaml:"url,omitempty"`
	// 邮箱
	Email string `yaml:"email,omitempty"`
}

type License struct {
	// 名称
	Name string `yaml:"name"`
	// 网址
	URL string `yaml:"url,omitempty"`
}
