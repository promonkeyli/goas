package goas

import (
	"errors"
	"fmt"

	"github.com/promonkeyli/goas/pkg/generater"
	"github.com/promonkeyli/goas/pkg/parser"
)

// Config Configuration for the generator
type Config struct {
	// Dirs Directories to scan for comments. e.g. ["./cmd", "./internal"]
	Dirs []string
	// Output Output directory path. e.g. "./api"
	Output string
}

// Run executes the parsing and generation process
func Run(cfg Config) error {
	if len(cfg.Dirs) == 0 {
		return errors.New("dirs cannot be empty")
	}
	if cfg.Output == "" {
		cfg.Output = "."
	}

	// Parse comments
	openapi, err := parser.Parse(cfg.Dirs)
	if err != nil {
		return fmt.Errorf("parse failed: %w", err)
	}

	// Generate file
	if err := generater.GenFiles(openapi, cfg.Output); err != nil {
		return fmt.Errorf("generate failed: %w", err)
	}

	return nil
}
