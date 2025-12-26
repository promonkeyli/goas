# GOAS 注释规范

### 全局注释

> **适用文件**：`main.go`
> **解析规则**：按行解析，参数间以空格分隔，`[]` 表示可选，`<>` 表示必填占位符。

| 分类 | 注解标记 | 必填 | 参数格式 | 说明 / 映射字段 | 示例 |
| :--: | :--- | :---: | :--- | :--- | :--- |
| **根配置** | **`@OpenAPI`** | **是** | `<version>` | OpenAPI 规范版本号<br>映射: `T.OpenAPI` | `// @OpenAPI 3.2.0` |
| | `@Self` | 否 | `<url>` | 文档自身的 URI<br>映射: `T.Self` | `// @Self https://api.com/doc.yaml` |
| | `@JsonSchemaDialect` | 否 | `<url>` | 默认 Json Schema 方案<br>映射: `T.JSONSchemaDialect` | `// @JsonSchemaDialect https://...` |
| **基本信息** | **`@Title`** | **是** | `<text>` | API 文档标题<br>映射: `Info.Title` | `// @Title 商城 API` |
| | **`@Version`** | **是** | `<string>` | API 业务版本号<br>映射: `Info.Version` | `// @Version 1.0.0` |
| | `@Summary` | 否 | `<text>` | API 简短摘要<br>映射: `Info.Summary` | `// @Summary 商城后端接口` |
| | `@Description` | 否 | `<markdown>` | API 详细描述 (支持多行)<br>映射: `Info.Description` | `// @Description ## 详情...` |
| | `@TermsOfService` | 否 | `<url>` | 服务条款链接<br>映射: `Info.TermsOfService` | `// @TermsOfService http://...` |
| **联系人** | `@Contact.Name` | 否 | `<string>` | 联系人姓名<br>>映射: `Info.Contact.Name` | `// @Contact.Name Support` |
| | `@Contact.Email` | 否 | `<email>` | 联系人邮箱<br>映射: `Info.Contact.Email` | `// @Contact.Email x@x.com` |
| | `@Contact.Url` | 否 | `<url>` | 联系人主页<br>映射: `Info.Contact.Url` | `// @Contact.Url http://x.com` |
| **许可证** | **`@License.Name`** | 否 | `<string>` | 许可证名称 (若填 License 则必填)<br>映射: `Info.License.Name` | `// @License.Name Apache 2.0` |
| | `@License.Identifier` | 否 | `<spdx_id>` | SPDX 许可证代码 (推荐)<br>映射: `Info.License.Identifier` | `// @License.Identifier Apache-2.0` |
| | `@License.Url` | 否 | `<url>` | 许可证 URL (与 ID 互斥)<br>映射: `Info.License.Url` | `// @License.Url http://...` |
| **服务列表** | **`@Server`** | 否 | `<url> [name=xxx] [desc]` | 定义服务器 (可重复)<br>映射: `T.Servers`<br>*注: name=前缀用于提取名称* | 1. `// @Server /v1 生产接口`<br>2. `// @Server /v1 name=prod 生产` |
| **外部文档** | `@ExternalDocs` | 否 | `<url> [desc]` | 全局外部文档<br>映射: `T.ExternalDocs` | `// @ExternalDocs http://wiki.com` |
| **标签** | **`@Tag.Name`** | 否 | `<string>` | **标签组开始**。标签名称。<br>映射: `Tag.Name` | `// @Tag.Name user` |
|  | `@Tag.Summary` | 否 | `<text>` | 标签短摘要 (3.2)<br>映射: `Tag.Summary` | `// @Tag.Summary 用户模块` |
| | `@Tag.Desc` | 否 | `<text>` | 标签详细说明。<br>映射: `Tag.Description` | `// @Tag.Desc 用户增删改查` |
| | `@Tag.Parent` | 否 | `<string>` | 父标签名称 (3.2)<br>映射: `Tag.Parent` | `// @Tag.Parent system` |
| | `@Tag.Kind` | 否 | `<string>` | 标签类型 (nav/badge等) (3.2)<br>映射: `Tag.Kind` | `// @Tag.Kind nav` |
| | `@Tag.Docs.Url` | 否 | `<url>` | 标签外部文档链接<br>映射: `Tag.ExternalDocs.URL` | `// @Tag.Docs.Url http://...` |
| | `@Tag.Docs.Desc` | 否 | `<text>` | 标签外部文档描述<br>映射: `Tag.ExternalDocs.Description` | `// @Tag.Docs.Desc 详情` |
| **安全方案** | **`@SecurityScheme`** | 否 | `<name> <type> [args...]` | 定义组件中的安全方案。<br>映射: `Components.SecuritySchemes`<br>**Type**: `apiKey`, `http`, `oauth2` | **ApiKey**: `// @SecurityScheme Key apiKey header Token`<br>**HTTP**: `// @SecurityScheme JWT http bearer JWT`<br>**OAuth2**: `// @SecurityScheme OAuth oauth2 implicit http://auth` |
| | `@SecurityScope` | 否 | `<scheme> <scope> <desc>` | 定义 OAuth2 的 Scope。<br>映射: `Flows.Scopes` | `// @SecurityScope OAuth write 读写` |
| **全局安全** | **`@Security`** | 否 | `<name> [scopes...]` | 应用全局安全限制。<br>映射: `T.Security` | `// @Security JWT` |

#### 示例

```go
package main

// @OpenAPI 3.2.0
// @Title Pet Store API
// @Version 1.0.0
// @Summary A sample pet store server
// @Description This is a Go implementation of the OpenAPI 3.2 Pet Store.
// @License.Name Apache 2.0
// @License.Identifier Apache-2.0
//
// @Server https://api.petstore.com/v1 name=prod Production
// @Server http://localhost:8080/v1 name=dev Local Dev
//
// @Tag.Name    pet
// @Tag.Summary 宠物管理
// @Tag.Kind    nav
//
// @SecurityScheme ApiKeyAuth apiKey header X-API-KEY
// @Security ApiKeyAuth
func main() {
    // ...
}
```

### 接口注释

> **适用范围**：API 接口处理函数（Handler Functions）上方。
> **解析规则**：按行解析，参数间以空格分隔，`[]` 表示可选，`<>` 表示必填占位符。

| 分类 | 注解标记 | 必填 | 参数格式 | 说明 / 映射字段 | 示例 |
| :--- | :--- | :---: | :--- | :--- | :--- |
| **路由配置** | **`@Router`** | **是** | `<path> [method]` | 定义路径和方法<br>映射: `Paths.{path}.{method}` | `// @Router /users/{id} [get]` |
| | `@Id` | 否 | `<string>` | 操作唯一标识符<br>映射: `Operation.OperationId` | `// @Id getUserById` |
| | `@Ignore` | 否 | - | 让工具忽略此函数，不生成文档。 | `// @Ignore` |
| | `@Deprecated` | 否 | - | 标记接口已废弃<br>映射: `Operation.Deprecated` | `// @Deprecated` |
| **基本信息** | **`@Summary`** | **是** | `<text>` | 接口简短摘要<br>映射: `Operation.Summary` | `// @Summary 获取用户详情` |
| | `@Description` | 否 | `<markdown>` | 接口详细描述 (支持多行)<br>映射: `Operation.Description` | `// @Description 查询用户的详细信息` |
| | `@Tags` | 否 | `<tag>[,tag...]` | 接口所属标签 (分组)<br>映射: `Operation.Tags` | `// @Tags user, admin` |
| **请求控制** | **`@Param`** | 否 | `<name> <in> <type> <req> <desc>` | 定义参数。<br>**in**: path/query/header/cookie/body/formData<br>**type**: string/int/file/struct<br>**req**: true/false<br>映射: `Parameters` 或 `RequestBody` | 1. `// @Param id path int true "用户ID"`<br>2. `// @Param q query string false "搜索"`<br>3. `// @Param req body model.User true "JSON"` |
| | `@Accept` | 否 | `<mime_type>` | 请求体类型 (Content-Type)<br>映射: `RequestBody.Content` Key | `// @Accept json,xml` |
| **响应定义** | **`@Success`** | **是*** | `<status> {<type>} <data> [desc]` | 成功响应 (*建议至少写一个)<br>**status**: 200/201...<br>**type**: object/array/string<br>**data**: Go类型或结构体路径 | `// @Success 200 {object} model.User "成功"`<br>`// @Success 200 {array} model.Item "列表"` |
| | `@Failure` | 否 | `<status> {<type>} <data> [desc]` | 失败响应<br>格式同上 | `// @Failure 400 {object} err.Resp "参数错误"` |
| | `@Produce` | 否 | `<mime_type>` | 响应类型 (Accept)<br>映射: `Responses.Content` Key | `// @Produce json` |
| | `@Header` | 否 | `<status> {<type>} <name> <desc>` | 响应头信息<br>映射: `Responses.Headers` | `// @Header 200 {string} Token "会话Token"` |
| **安全与扩展** | `@Security` | 否 | `<name> [scopes...]` | 覆盖全局安全设置<br>映射: `Operation.Security` | `// @Security ApiKeyAuth` |
| | `@ExternalDocs` | 否 | `<url> [desc]` | 接口级外部文档<br>映射: `Operation.ExternalDocs` | `// @ExternalDocs http://wiki.com 详情` |

#### 示例

```go
package user

// GetUser 获取用户详情
// @Summary 获取用户详情
// @Description 根据用户 ID 查询用户的详细信息，包括角色和权限。
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户 ID"
// @Param type query string false "视图类型"
// @Success 200 {object} model.UserResponse "用户信息"
// @Failure 400 {object} model.ErrorResponse "ID 无效"
// @Failure 404 {object} model.ErrorResponse "用户不存在"
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
    // ...
}
```