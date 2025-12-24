# Goas 静态代码分析工具开发计划
既然你已经选定了 **Goas** 这个极简的名字，并且明确了“仅支持 OpenAPI 3.0+”以及“轻量化”的目标，这份开发路线图将分为五个阶段，从核心引擎到高级特性，帮助你逐步构建出一款具备竞争力的工具。

---

### 🚀 Goas 开发路线图 (Development Roadmap)

#### 第一阶段：核心骨架与扫描 (Foundation)
**目标**：建立项目基础，实现对 Go 代码的初步读取和注释提取。

*   **1.1 环境初始化**：
    *   使用 `cobra` 初始化 CLI 框架（实现 `goas init` 和 `goas gen` 命令）。
    *   选定 OpenAPI 3 模型库（推荐 `github.com/getkin/kin-openapi/openapi3`）。
*   **1.2 静态分析扫描器**：
    *   使用 `golang.org/x/tools/go/packages` 加载目标项目（比 `go/parser` 更强，支持跨包分析）。
    *   遍历 AST，定位所有的函数声明 (`ast.FuncDecl`) 及其关联的文档注释 (`Doc`)。
*   **1.3 极简 DSL 定义**：
    *   实现对 `@Router` 和 `@Summary` 的基础正则解析。
    *   **目标产出**：能扫描项目并打印出所有带注释的路由列表。

#### 第二阶段：深度解析引擎 (The "Brain")
**目标**：攻克最难的技术点——将 Go 结构体精准转化为 OpenAPI Schema。

*   **2.1 类型提取器 (Type Extractor)**：
    *   解析函数注释中的请求/响应类型（如 `@Response 200 {object} User`）。
    *   利用 `go/types` 寻找该结构体的真实定义（处理跨包引用）。
*   **2.2 Schema 转换逻辑**：
    *   实现基础类型映射（`string` -> `string`, `int64` -> `integer/int64` 等）。
    *   **JSON Tag 处理**：解析 `json:"name,omitempty"`，确定字段名和是否必填。
    *   **递归解析**：处理结构体嵌套、切片（Array）和 Map。
*   **2.3 泛型支持 (Go 1.18+)**：
    *   识别 `ast.IndexExpr`，解析如 `Response[User]` 这种带参数的类型。

#### 第三阶段：组装与生成 (Assembly)
**目标**：生成一份符合规范的、可在 Swagger UI 或 Apifox 中打开的 JSON/YAML。

*   **3.1 Components 管理**：
    *   实现全局 `components/schemas` 的收集，确保相同的结构体只定义一次，并在接口处使用 `$ref` 引用。
*   **3.2 文档组装**：
    *   填充 `info`, `paths`, `servers`, `security` 等 OpenAPI 核心字段。
*   **3.3 多格式导出**：
    *   支持输出为 `openapi.json` 或 `openapi.yaml`。
*   **目标产出**：一个完整的、可验证的 OpenAPI 3.0 离线文件。

#### 第四阶段：增强特性 (The "Edge")
**目标**：让 Goas 比现有工具更智能，减少手动注释。

*   **4.1 自动验证支持 (Validation Tags)**：
    *   解析 `validate:"required,min=10,max=100"` 等标签，自动转化为 OpenAPI 的 `minLength`, `maximum` 等约束。
*   **4.2 路由自动推导 (选做)**：
    *   针对 Gin/Echo 等常用框架，尝试直接从代码行 `r.GET("/ping", handler)` 中推导路径，进一步减少注释压力。
*   **4.3 多 Server 配置**：
    *   支持在全局注释或配置文件中定义不同的 API 运行环境。

#### 第五阶段：工程化与生态 (Ecosystem)
**目标**：提升开发者体验，让项目具备生产力。

*   **5.1 热重载模式 (Watch Mode)**：
    *   实现 `goas gen --watch`，监听文件变动自动更新文档。
*   **5.2 文档校验 (Linter)**：
    *   在生成后自动运行规范校验，提示缺少的描述或错误的引用。
*   **5.3 CI/CD 集成**：
    *   提供 GitHub Action 示例，实现代码合并自动推送文档到 Apifox/Postman。

---

### 🛠 关键技术栈推荐

| 功能模块 | 推荐库 | 理由 |
| :--- | :--- | :--- |
| **CLI 框架** | `spf13/cobra` | Go CLI 事实上的标准，功能强大且稳定。 |
| **项目扫描** | `golang.org/x/tools/go/packages` | 官方维护，能完美处理模块化后的跨包类型查找。 |
| **OpenAPI 模型** | `getkin/kin-openapi` | 目前社区最成熟的 OpenAPI 3 读写校验库。 |
| **配置读取** | `spf13/viper` | 配合 Cobra 使用，支持 YAML/JSON/Env 配置。 |

---

### 💡 开发者建议：第一个版本 (v0.1.0) 应该关注什么？

不要试图第一天就支持所有复杂的泛型和深层嵌套。你的 **MVP (最小可行性产品)** 应该是：
1.  **能够解析当前包下的简单结构体**。
2.  **支持 `@Summary`, `@Router`, `@Response` 三大核心标签**。
3.  **生成一个能被 Swagger UI 成功渲染的 JSON 文件**。

**Goas** 的核心卖点是**“轻”**和**“准”**。只要你能比 `swag` 在 OpenAPI 3 格式上生成得更标准，开发者就愿意迁移。

你需要我针对 **“如何递归解析结构体字段并生成 JSON Schema”** 给你提供一个核心算法的伪代码吗？这是开发过程中最容易卡住的地方。