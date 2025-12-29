package main

import (
	"flag"
	"log/slog"
	"os"
	"strings"

	"github.com/promonkeyli/goas/pkg/generater"
	"github.com/promonkeyli/goas/pkg/parser"
)

func main() {
	// 1. 变量定义： 扫描的目录或者文件(多个目录使用逗号分隔的字符串： "./a,./b")/输出文件路径
	var dir, output string

	// 2. 变量绑定
	flag.StringVar(&dir, "dir", "", "扫描的目录，多个目录使用逗号分隔")
	flag.StringVar(&output, "output", "./api", "输出文件路径")

	// 3. 解析命令行参数
	flag.Parse()

	// 4.参数校验
	if dir == "" {
		slog.Error("参数错误: 需要指定扫描目录或者文件路径")
		os.Exit(1)
	}

	// 5. 逗号分隔参数,调用解析器
	var dirs []string
	for _, s := range strings.Split(dir, ",") {
		s = strings.TrimSpace(s) // 去掉空格
		if s != "" {
			dirs = append(dirs, s)
		}
	}
	// 调用解析器
	openapi, err := parser.Parse(dirs)
	if err != nil {
		slog.Error("解析失败", "error", err)
		os.Exit(1)
	}

	// 6. 输出文件
	if err := generater.GenFiles(openapi, output); err != nil {
		slog.Error("生成文件失败", "error", err)
		os.Exit(1)
	}
}
