package parser

import (
	"fmt"

	"github.com/promonkeyli/goas/pkg/model"
	"golang.org/x/tools/go/packages"
)

func Parse(dirs []string) (*model.T, error) {
	fmt.Printf("开始扫描目录: %v\n", dirs)

	// 1. 初始化
	p := newProcessor()

	// 1. 配置加载模式
	// 这是 go/packages 最强大的地方，我们需要：
	// - NeedName: 包名
	// - NeedFiles: 文件路径
	// - NeedSyntax: AST 语法树 (为了解析注释)
	// - NeedTypes: 类型信息 (为了解析结构体字段类型)
	// - NeedTypesInfo: 具体的 AST 到 Type 的映射
	// - NeedImports: 解析 import 关系
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles |
			packages.NeedSyntax | packages.NeedTypes |
			packages.NeedTypesInfo | packages.NeedImports,
		Tests: false, // 通常不需要扫描测试文件
	}

	// 2. 执行加载
	// packages.Load 支持变长参数，我们直接把 dirs 切片展开传进去
	pkgs, err := packages.Load(cfg, dirs...)
	if err != nil {
		return nil, fmt.Errorf("加载包失败: %w", err)
	}

	// 3. 错误检查
	// packages.Load 即使有语法错误也可能返回 err=nil，需要检查返回的包里是否有错误
	if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("源码中存在错误，无法继续解析")
	}

	// 4. 【建立索引】构建包映射表
	// 这步至关重要：把 slice 转成 map，后续查 "github.com/lib/pq" 这种路径时能 O(1) 找到
	for _, pkg := range pkgs {
		p.addToIndex(pkg)
	}

	// 5. 【执行扫描】
	// 遍历所有请求的包，寻找 Controller/API 定义
	for _, pkg := range pkgs {
		p.scanPackage(pkg)
	}

	return p.OpenAPI, nil
}
