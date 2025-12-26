# OpenAPI 3.2.0 字段参考文档

本文档总结了 OpenAPI 3.2.0 规范中的核心对象及其字段，提供详尽的中文描述，方便开发人员查阅。

1. [根对象 (OpenAPI Object)](#1-根对象-openapi-object)
2. [元数据对象 (Info Object)](#2-元数据对象-info-object)
3. [服务器对象 (Server Object)](#3-服务器对象-server-object)
4. [组件对象 (Components Object)](#4-组件对象-components-object)
5. [路径对象 (Paths Object)](#5-路径对象-paths-object)
6. [操作对象 (Operation Object)](#6-操作对象-operation-object)
7. [参数对象 (Parameter Object)](#7-参数对象-parameter-object)
8. [请求体对象 (Request Body Object)](#8-请求体对象-request-body-object)
9. [响应对象 (Responses Object)](#9-响应对象-responses-object)
10. [数据模型对象 (Schema Object)](#10-数据模型对象-schema-object)
11. [安全对象 (Security Object)](#11-安全对象-security-object)
12. [其他辅助对象](#12-其他辅助对象)

---

## 1. 根对象 (OpenAPI Object)
API 文档的根节点。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **openapi** | `string` | 是 | **REQUIRED**. OpenAPI 规范版本号（如 `3.2.0`）。 |
| **info** | [Info](#2-元数据对象-info-object) | 是 | **REQUIRED**. API 的元数据信息。 |
| **jsonSchemaDialect** | `string` | 否 | 用于文档中所有 Schema 对象的默认方言。 |
| **servers** | [`[]Server`](#3-服务器对象-server-object) | 否 | 提供的服务器连接信息数组。 |
| **paths** | [Paths](#5-路径对象-paths-object) | 否 | API 的可用路径及操作。 |
| **webhooks** | `map[string]PathItem` | 否 | 描述受入站 HTTP 回调。 |
| **components** | [Components](#4-组件对象-components-object) | 否 | 包含各种可重用的对象组件。 |
| **security** | `[]SecurityRequirement` | 否 | API 的全局安全要求。 |
| **tags** | `[]Tag` | 否 | 标记列表，用于对 API 进行分类。 |
| **externalDocs** | `ExternalDocs` | 否 | 外部扩展文档。 |
| **$self** | `string` | 否 | 当前文档的 base URI。 |

---

## 2. 元数据对象 (Info Object)
提供 API 的基本信息。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **title** | `string` | 是 | **REQUIRED**. API 的标题。 |
| **summary** | `string` | 否 | API 的简短摘要。 |
| **description** | `string` | 否 | API 的详细描述（支持 CommonMark）。 |
| **termsOfService** | `string` | 否 | 服务条款的 URL。 |
| **contact** | `Contact` | 否 | 联系信息。 |
| **license** | `License` | 否 | 许可证信息。 |
| **version** | `string` | 是 | **REQUIRED**. API 自身的版本号。 |

---

## 3. 服务器对象 (Server Object)
描述目标主机的连接信息。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **url** | `string` | 是 | **REQUIRED**. 目标主机的 URL。 |
| **description** | `string` | 否 | URL 的描述。 |
| **name** | `string` | 否 | 变量替换中使用的唯一引用名称。 |
| **variables** | `map[string]ServerVariable` | 否 | URL 模板替换的变量映射。 |

### Server Variable Object
| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **enum** | `[]string` | 否 | 变量的可选值枚举。 |
| **default** | `string` | 是 | **REQUIRED**. 默认值。 |
| **description** | `string` | 否 | 变量的详细描述。 |

---

## 4. 组件对象 (Components Object)
用于存放各种可重用的 OpenAPI 对象。

| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **schemas** | `map[string]Schema` | 可重用的数据模型。 |
| **responses** | `map[string]Response` | 可重用的响应定义。 |
| **parameters** | `map[string]Parameter` | 可重用的参数定义。 |
| **examples** | `map[string]Example` | 可重用的示例。 |
| **requestBodies** | `map[string]RequestBody` | 可重用的请求体定义。 |
| **headers** | `map[string]Header` | 可重用的头部定义。 |
| **securitySchemes** | `map[string]SecurityScheme` | 安全方案定义。 |
| **links** | `map[string]Link` | 可重用的链接定义。 |
| **callbacks** | `map[string]Callback` | 可重用的回调定义。 |
| **pathItems** | `map[string]PathItem` | 可重用的路径项。 |

---

## 5. 路径对象 (Paths Object)
定义 API 的各个路径。

### Path Item Object
| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **$ref** | `string` | 允许引用另一个路径项。 |
| **summary** | `string` | 简短摘要。 |
| **description** | `string` | 详细描述。 |
| **get/put/post...** | [Operation](#6-操作对象-operation-object) | 对应的 HTTP 方法操作定义。 |
| **servers** | `[]Server` | 该路径特有的服务器覆盖。 |
| **parameters** | `[]Parameter` | 该路径下所有操作共享的参数。 |

---

## 6. 操作对象 (Operation Object)
描述单个 HTTP 方法的 API 操作。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **tags** | `[]string` | 否 | 用于分类的标签列表。 |
| **summary** | `string` | 否 | 简短的操作摘要。 |
| **description** | `string` | 否 | 详细的功能描述。 |
| **operationId** | `string` | 否 | 唯一的操作 ID。 |
| **parameters** | `[]Parameter` | 否 | 操作参数列表。 |
| **requestBody** | [RequestBody](#8-请求体对象-request-body-object) | 否 | 请求体内容定义。 |
| **responses** | [Responses](#9-响应对象-responses-object) | 是 | **REQUIRED**. 各种可能的响应定义。 |
| **callbacks** | `map[string]Callback` | 否 | 相关的回调映射。 |
| **deprecated** | `bool` | 否 | 是否已弃用。 |
| **security** | `[]SecurityRequirement` | 否 | 操作特定的安全要求覆盖。 |

---

## 7. 参数对象 (Parameter Object)
定义操作的参数（Query, Header, Path, Cookie）。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **name** | `string` | 是 | **REQUIRED**. 参数名称。 |
| **in** | `string` | 是 | **REQUIRED**. 位置 (`query`, `querystring`, `header`, `path`, `cookie`)。 |
| **description** | `string` | 否 | 描述参数的用途。 |
| **required** | `bool` | 否 | 是否必填（`path` 参数必须为 `true`）。 |
| **deprecated** | `bool` | 否 | 是否已弃用。 |
| **schema** | [Schema](#10-数据模型对象-schema-object) | 否 | 参数的数据结构。 |
| **style** | `string` | 否 | 序列化风格（如 `simple`, `form` 等）。 |
| **explode** | `bool` | 否 | 对于数组/对象，是否展开。 |
| **example/examples** | `any` | 否 | 示例值。 |

---

## 8. 请求体对象 (Request Body Object)
描述请求中发送的内容。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **description** | `string` | 否 | 请求体的描述。 |
| **content** | `map[string]MediaType` | 是 | **REQUIRED**. 媒体类型及其 Schema 的映射。 |
| **required** | `bool` | 否 | 请求体是否必选。 |

### Media Type Object
| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **schema** | [Schema](#10-数据模型对象-schema-object) | 描述数据的模型。 |
| **itemSchema** | [Schema](#10-数据模型对象-schema-object) | (3.2.0 新增) 针对流式数组每一项的 Schema。 |
| **example(s)** | `any` | 对应的媒体类型示例。 |
| **encoding** | `map[string]Encoding` | 对象属性的特殊编码定义。 |

---

## 9. 响应对象 (Responses Object)
描述操作可能返回的所有响应。

| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **default** | [Response](#response-object) | 否 | 默认响应。 |
| **[HTTPCode]** | [Response](#response-object) | 否 | HTTP 状态码（如 `200`）对应的响应。 |

### Response Object
| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **description** | `string` | 是 | **REQUIRED**. 响应内容的简要描述。 |
| **summary** | `string` | 否 | (3.2.0 新增) 响应的简短摘要。 |
| **headers** | `map[string]Header` | 否 | 响应头列表。 |
| **content** | `map[string]MediaType` | 否 | 响应体内容的媒体类型映射。 |
| **links** | `map[string]Link` | 否 | 运行时链接定义。 |

---

## 10. 数据模型对象 (Schema Object)
基于 JSON Schema 2020-12 的核心模型。

| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **$id / $ref** | `string` | 唯一标识符 / 引用。 |
| **type** | `any` | 数据类型 (`string`, `number`, `object`, `array`, `boolean`, `null`)，可为数组。 |
| **format** | `string` | 数据的格式（如 `date-time`, `int64`）。 |
| **title/description** | `string` | 元数据说明。 |
| **enum** | `[]any` | 枚举值。 |
| **const** | `any` | (3.2.0 新增) 常量值。 |
| **properties** | `map[string]Schema` | 对象的属性定义。 |
| **items** | [Schema](#schema-object) | 数组项的定义。 |
| **prefixItems** | `[]Schema` | (3.2.0 新增) 数组前缀项定义（元组）。 |
| **required** | `[]string` | 对象的必选属性列表。 |
| **allOf / anyOf / oneOf** | `[]Schema` | 组合模式。 |
| **maximum / minimum** | `any` | 数值范围限制。 |
| **maxLength / minLength** | `int` | 字符串长度限制。 |
| **maxItems / minItems** | `int` | 数组项数限制。 |
| **contentMediaType**| `string` | (3.2.0 新增) 字符串内容的媒体类型。 |
| **contentSchema** | [Schema](#schema-object) | (3.2.0 新增) 字符串内容的具体 Schema。 |
| **nullable** | `bool` | 是否允许为 null（OpenAPI 特定，3.1+ 建议用 type 包含 null）。 |
| **discriminator** | `Discriminator`| 多态鉴别器。 |
| **xml** | `XML` | XML 序列化配置。 |

---

## 11. 安全对象 (Security Object)

### Security Scheme Object
| 字段名 | 类型 | 必选 | 描述 |
| :--- | :--- | :---: | :--- |
| **type** | `string` | 是 | **REQUIRED**. 类型：`apiKey`, `http`, `mutualTLS`, `oauth2`, `openIdConnect`。 |
| **name** | `string` | (apiKey 是) | header, query 或 cookie 中使用的参数名。 |
| **in** | `string` | (apiKey 是) | 位置：`query`, `header`, `cookie`。 |
| **scheme** | `string` | (http 是) | HTTP 认证方案（如 `basic`, `bearer`）。 |
| **flows** | `OAuthFlows` | (oauth2 是) | OAuth 流程配置。 |
| **openIdConnectUrl**| `string` | (oidc 是) | OIDC 配置 URL。 |

---

## 12. 其他辅助对象

### Tag Object
| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **name** | `string` | **REQUIRED**. 标签名。 |
| **summary** | `string` | (3.2.0 新增) 标签摘要。 |
| **description** | `string` | 标签详情。 |
| **parent** | `string` | (3.2.0 新增) 父标签名。 |
| **kind** | `string` | (3.2.0 新增) 标签分类（如 `nav`, `audience`）。 |

### Example Object
| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **summary** | `string` | 示例摘要。 |
| **description** | `string` | 示例详细描述。 |
| **dataValue** | `any` | (3.2.0 新增) 符合 Schema 的原始数据示例。 |
| **serializedValue**| `string` | (3.2.0 新增) 序列化后的字符串示例。 |
| **value** | `any` | (已弃用) 嵌入的示例值。 |
| **externalValue** | `string` | 外部示例的 URI。 |

---

### Reference Object
| 字段名 | 类型 | 描述 |
| :--- | :--- | :--- |
| **$ref** | `string` | **REQUIRED**. 引用的 URI。 |
| **summary** | `string` | 对被引用对象的摘要覆盖。 |
| **description** | `string` | 对被引用对象的描述覆盖。 |

---
*本文档由 Antigravity 整理完成，基于 OpenAPI 3.2.0 官方规范。*
