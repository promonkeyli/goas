package internal

import (
	"fmt"
)

// GetUser 获取用户详情
// @Summary 获取用户详情
// @Description 根据用户 ID 查询用户的详细信息
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} model.Response[model.User] "成功"
// @Failure 400 {object} model.ErrorResponse "参数错误"
// @Failure 404 {object} model.ErrorResponse "用户不存在"
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
// @Success 200 {object} model.Response[model.PageList[model.User]] "成功"
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
// @Param user body model.User true "用户信息"
// @Success 201 {object} model.Response[model.User] "创建成功"
// @Failure 400 {object} model.ErrorResponse "参数错误"
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
// @Param user body model.User true "用户信息"
// @Success 200 {object} model.Response[model.User] "更新成功"
// @Failure 400 {object} model.ErrorResponse "参数错误"
// @Failure 404 {object} model.ErrorResponse "用户不存在"
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
// @Failure 404 {object} model.ErrorResponse "用户不存在"
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
// @Success 200 {object} model.Response[string] "上传成功"
// @Failure 400 {object} model.ErrorResponse "上传失败"
// @Router /users/{id}/avatar [post]
func UploadAvatar() {
	fmt.Println("upload avatar handler")
}
