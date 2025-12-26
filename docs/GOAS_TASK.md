# GOAS 开发计划书
```markdown
**项目名称**: Goas (Go OpenAPI Static analysis)
**目标版本**: OpenAPI 3.2.0
**开发语言**: Go (Golang)
**核心目标**: 一个轻量级、非侵入式的静态代码分析工具，用于从 Go 代码注释生成 OpenAPI 3.2 文档。

## 1. 🏗 系统架构设计 (System Architecture)

### 1.1 推荐目录结构
遵循 Go 标准项目布局 (Standard Go Project Layout)：

```text
goas/
├── cmd/
│   └── goas/           # 程序入口 (CLI)
│       └── main.go
├── pkg/
│   ├── model/          # OpenAPI 3.2 数据模型 (T, Info, Schema 等)
│   ├── analysis/       # AST 解析与注释提取核心
│   ├── schema/         # 类型转换引擎 (Go Type -> OpenAPI Schema)
│   ├── generator/      # 最终产物生成 (YAML/JSON)
│   └── config/         # 配置处理
├── internal/
│   └── utils/          # 工具函数 (Tag 解析, 字符串处理)
├── testdata/           # 用于测试的 Go 代码样例
├── go.mod
└── README.md
```

### 1.2 核心处理流程 (Pipeline)
工具运行分为四个明确的阶段：
1.  **扫描 (Scanning)**: 使用 `golang.org/x/tools/go/packages` 加载 Go 包信息。
2.  **解析 (Parsing)**: 解析 `main.go` (全局配置) 和 Handler 函数 (接口操作) 中的注释。
3.  **提取 (Extracting)**: 将 Go 的结构体/类型转换为 OpenAPI Schemas (处理 JSON tag、递归引用等)。
4.  **生成 (Generating)**: 将结果序列化为 `openapi.yaml` 文件。

---

## 📅 2. 分阶段实施路线图 (Development Roadmap)

### 第一阶段：基础设施与全局配置
**目标**：初始化 CLI 并生成一份有效（包含头部信息但无路径）的 OpenAPI 文档。

- [ ] **1.1 项目初始化**
    - [ ] 执行 `go mod init`。
    - [ ] 安装依赖：`spf13/cobra` (CLI), `gopkg.in/yaml.v3`。
    - [ ] 创建 `goas` 命令结构 (`root`, `init`, `gen`)。
- [ ] **1.2 模型层实现 (`pkg/model`)**
    - [ ] 基于 OpenAPI 3.2 定义 `T`, `Info`, `Server`, `Tag`, `ExternalDocs` 等结构体。
    - [ ] 为所有结构体添加 `yaml:"..."` 标签。
- [ ] **1.3 Main 包解析器 (`pkg/analysis`)**
    - [ ] 实现 `Loader` 以解析 `main.go`。
    - [ ] 实现正则/逻辑以解析全局注释：
        - [ ] `@OpenAPI`, `@Title`, `@Version`
        - [ ] `@Server`
        - [ ] `@Tag` (全局定义)
        - [ ] `@SecurityScheme` (组件定义)
    - [ ] **产出物**：一个包含 `openapi`, `info`, `servers` 字段的 YAML 文件。

### 第二阶段：路由与基础操作
**目标**：扫描所有 Handler 函数并生成 API 路径骨架。

- [ ] **2.1 函数扫描器**
    - [ ] 遍历 AST，寻找带有 `// @Router` 注释的函数声明。
    - [ ] 解析 HTTP Method 和 Path，映射到 `T.Paths`。
- [ ] **2.2 操作元数据 (Operation Metadata)**
    - [ ] 解析 `@Summary`, `@Description`, `@Id`, `@Deprecated`。
    - [ ] 解析 `@Tags` 并关联到全局 Tag。
    - [ ] 解析 `@Ignore` 以跳过特定函数。
- [ ] **2.3 基础参数解析**
    - [ ] 解析 `@Param`，支持 `path`, `query`, `header` 位置。
    - [ ] **限制**：暂时仅支持基础类型 (`int`, `string`, `bool`)。

### 第三阶段：类型系统与 Schema 引擎 (核心难点)
**目标**：将 Go 结构体转换为 JSON Schemas 并处理组件注册。

- [ ] **3.1 深度类型分析**
    - [ ] 配置 `go/packages` 开启 `NeedTypes | NeedTypesInfo | NeedSyntax`。
- [ ] **3.2 Schema 转换器 (`pkg/schema`)**
    - [ ] 实现 `TypeToSchema(t types.Type) *model.Schema`。
    - [ ] **基础类型**：映射 `int`->`integer`, `float64`->`number`, `time.Time`->`string(date-time)`。
    - [ ] **切片/Map**：递归解析 `[]T` 和 `map[string]T`。
    - [ ] **结构体 (Structs)**：
        - [ ] 遍历字段。
        - [ ] 读取 `json` tag 确定字段名。
        - [ ] 读取 `validate` tag 确定 `required` 字段。
        - [ ] 读取字段后的行内注释作为 `description`。
- [ ] **3.3 组件注册中心 (Component Registry)**
    - [ ] 实现缓存以存储已解析的结构体。
    - [ ] 检测命名结构体并注册到 `T.Components.Schemas`。
    - [ ] 将内联定义替换为引用 `$ref: "#/components/schemas/Name"`。
- [ ] **3.4 请求与响应集成**
    - [ ] 升级 `@Param ... body` 解析逻辑，调用 Schema 引擎。
    - [ ] 升级 `@Success` / `@Failure` 解析逻辑，调用 Schema 引擎。

### 第四阶段：安全与高级特性
**目标**：提供完整的 OpenAPI 3.2 支持。

- [ ] **4.1 安全特性**
    - [ ] 解析 `@SecurityScheme` 注释到 `Components.SecuritySchemes`。
    - [ ] 解析 Handler 上的 `@Security` 注释到 `Operation.Security`。
- [ ] **4.2 高级类型支持**
    - [ ] 支持 `multipart/form-data` (文件上传)。
    - [ ] 支持 Go 1.18+ 泛型 (例如 `Response[User]`)。
    - [ ] 支持匿名结构体。
- [ ] **4.3 OpenAPI 3.2 合规性**
    - [ ] 确保使用 `type: ["string", "null"]` 代替 `nullable: true`。
    - [ ] 验证 `jsonSchemaDialect` 的输出。

### 第五阶段：打磨与发布
**目标**：生产就绪的工具。

- [ ] **5.1 格式化**
    - [ ] 确保 YAML 输出顺序符合阅读习惯 (Info -> Servers -> Paths -> Components)。
- [ ] **5.2 测试**
    - [ ] 在 `testdata/` 中创建复杂的 Go 代码示例。
    - [ ] 为解析器和 Schema 转换器编写单元测试。
- [ ] **5.3 文档**
    - [ ] 编写 `README.md`，包含安装指南。
    - [ ] 包含注释规范速查表 (Annotation Spec Tables)。

---

## 🛠 3. 技术栈选型 (Tech Stack)

| 模块 | 库/工具 | 用途 |
| :--- | :--- | :--- |
| **CLI** | `github.com/spf13/cobra` | 构建命令行交互界面 |
| **AST/Parser** | `golang.org/x/tools/go/packages` | 加载 Go 包、类型信息和 AST |
| **YAML** | `gopkg.in/yaml.v3` | 将最终的 OpenAPI 结构体序列化为 YAML |
| **Testing** | `github.com/stretchr/testify` | 单元测试断言库 |

---

## 📝 4. 关键数据结构回顾

### 4.1 全局上下文 (Global Context)
用于在解析的各个阶段传递状态数据。

```go
type ParserContext struct {
    Pkg          *packages.Package        // 当前包信息
    Doc          *model.T                 // 正在构建的根 OpenAPI 文档对象
    
    // Schema 缓存，处理 $ref 引用和递归
    // Key: 完全限定类型名 (例如: "github.com/my/project/model.User")
    SchemaCache  map[string]*model.Schema 
}
```

### 4.2 注释规范速查 (Quick Reference)

#### Main.go (全局配置)
| 注解 | 示例 |
| :--- | :--- |
| `@OpenAPI` | `// @OpenAPI 3.2.0` |
| `@Title` | `// @Title 商城 API` |
| `@Server` | `// @Server /v1 name=prod 生产环境` |
| `@Tag.Name` | `// @Tag.Name user` |
| `@SecurityScheme` | `// @SecurityScheme MyKey apiKey header X-Token` |

#### Handler Func (接口操作)
| 注解 | 示例 |
| :--- | :--- |
| `@Router` | `// @Router /users/{id} [get]` |
| `@Summary` | `// @Summary 获取用户` |
| `@Param` | `// @Param id path int true "用户ID"` |
| `@Success` | `// @Success 200 {object} model.User` |
| `@Security` | `// @Security MyKey` |

---

## 🚀 5. 快速开始 (开发者指南)

1. **克隆仓库**: `git clone ...`
2. **创建模型**: 将 `T` 结构体定义复制到 `pkg/model/openapi.go`。
3. **运行加载器**: 编写一个简单的 `main` 函数，使用 `go/packages` 打印示例项目的 AST，验证环境。
```