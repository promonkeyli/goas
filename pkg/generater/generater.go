package generater

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/promonkeyli/goas/pkg/model"
)

func GenFiles(openAPI *model.T, outputPath string) error {
	// 1. 确保输出目录存在
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		// 创建目录，权限 0755
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("创建目录失败: %w", err)
		}
	}

	// 2. 生成 JSON 内容
	// 使用 MarshalIndent 可以在输出时进行格式化（带缩进），方便阅读
	jsonBytes, err := json.MarshalIndent(openAPI, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %w", err)
	}

	// 3. 定义文件路径
	jsonPath := filepath.Join(outputPath, "openapi.json")

	// 4. 写入 JSON 文件
	// 0644 代表文件所有者可读写，其他人可读
	if err := os.WriteFile(jsonPath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("写入 JSON 文件失败: %w", err)
	}
	fmt.Printf("生成成功: %s\n", jsonPath)

	return nil
}
