# Go OpenAPI 3.2 Global Annotation Specification

**适用文件**：`main.go`
**解析位置**：`package main` 之前或 `main()` 函数之前。
**解析规则**：按行解析，参数间以空格分隔，`[]` 表示可选，`<>` 表示必填占位符。

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

---

### 使用示例

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