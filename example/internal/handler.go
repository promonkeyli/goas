package internal

import "fmt"

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

// GetUser 获取用户详情
// @Summary 获取用户详情
// @Description 根据用户 ID 查询用户的详细信息
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} Response[User] "成功"
// @Failure 400 {object} ErrorResponse "参数错误"
// @Failure 404 {object} ErrorResponse "用户不存在"
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser() {
	fmt.Println("get user handler")
}

// ListUsers 获取用户列表
// @Summary 获取用户列表
// @Description 分页查询用户列表
// @Tags user
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "搜索关键词"
// @Success 200 {object} Response[PageList[User]] "成功"
// @Failure 400 {object} ErrorResponse "参数错误"
// @Router /users [get]
func ListUsers() {
	fmt.Println("list users handler")
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建一个新用户
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "用户信息"
// @Success 201 {object} Response[User] "创建成功"
// @Failure 400 {object} ErrorResponse "参数错误"
// @Router /users [post]
func CreateUser() {
	fmt.Println("create user handler")
}

// UpdateUser 更新用户
// @Summary 更新用户信息
// @Description 更新指定用户的信息
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body User true "用户信息"
// @Success 200 {object} Response[User] "更新成功"
// @Failure 400 {object} ErrorResponse "参数错误"
// @Failure 404 {object} ErrorResponse "用户不存在"
// @Router /users/{id} [put]
func UpdateUser() {
	fmt.Println("update user handler")
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除指定用户
// @Tags user
// @Produce json
// @Param id path int true "用户ID"
// @Success 204 "删除成功"
// @Failure 404 {object} ErrorResponse "用户不存在"
// @Router /users/{id} [delete]
// @Deprecated
func DeleteUser() {
	fmt.Println("delete user handler")
}

// UploadAvatar 上传头像
// @Summary 上传用户头像
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "用户ID"
// @Param avatar formData file true "头像文件"
// @Success 200 {object} Response[string] "上传成功"
// @Failure 400 {object} ErrorResponse "上传失败"
// @Router /users/{id}/avatar [post]
func UploadAvatar() {
	fmt.Println("upload avatar handler")
}
