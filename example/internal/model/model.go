package model

// User 用户模型
type User struct {
	ID       int    `json:"id" desc:"用户ID"`
	Name     string `json:"name" desc:"用户名"`
	Email    string `json:"email,omitempty" desc:"邮箱"`
	Age      int    `json:"age,omitempty" desc:"年龄"`
	IsActive bool   `json:"is_active" desc:"是否激活"`
}

// Response 通用响应结构 (泛型)
type Response[T any] struct {
	Code    int    `json:"code" desc:"状态码"`
	Message string `json:"message" desc:"消息"`
	Data    T      `json:"data" desc:"数据"`
}

// PageList 分页列表 (泛型)
type PageList[T any] struct {
	Total int `json:"total" desc:"总数"`
	Page  int `json:"page" desc:"当前页"`
	Size  int `json:"size" desc:"每页数量"`
	Items []T `json:"items" desc:"数据列表"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code" desc:"错误码"`
	Message string `json:"message" desc:"错误信息"`
	Details string `json:"details,omitempty" desc:"详细信息"`
}

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" desc:"用户名"`
	Password string `json:"password" desc:"密码"`
}

// LoginRes 登录详情
type LoginRes struct {
	AccessToken  string `json:"access_token" desc:"访问凭证"`
	RefreshToken string `json:"refresh_token" desc:"刷新凭证"`
}
