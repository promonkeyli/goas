package main

import "fmt"

// @OpenAPI 3.2.0
// @Title Pet Store API
// @Version 1.0.0
// @Summary A sample pet store server
// @Description This is a Go implementation of the OpenAPI 3.2 Pet Store.
// @Description Supports user management and file uploads.
// @TermsOfService http://localhost:8080/terms
//
// @Contact.Name API Support
// @Contact.URL http://localhost:8080/support
// @Contact.Email support@example.com
//
// @License.Name Apache 2.0
// @License.Identifier Apache-2.0
//
// @Server http://localhost:8080/v1 name=dev 开发环境
// @Server https://api.example.com/v1 name=prod 生产环境
//
// @Tag.Name user
// @Tag.Summary 用户管理
// @Tag.Desc 用户的增删改查操作
// @Tag.Kind nav
//
// @Tag.Name admin
// @Tag.Summary 管理后台
// @Tag.Parent user
//
// @SecurityScheme ApiKeyAuth apiKey header X-API-KEY
// @Security ApiKeyAuth
func main() {
	fmt.Println("cmd file")
}
