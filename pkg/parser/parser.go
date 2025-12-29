package parser

import (
	"fmt"

	"github.com/promonkeyli/goas/pkg/model"
)

func Parse(dir []string) (*model.T, error) {
	fmt.Printf("扫描目录: %v", dir)

	// 1. 结构初始化
	openapi := &model.T{}

	// 返回数据
	return openapi, nil
}
