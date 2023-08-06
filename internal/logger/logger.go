package logger

import (
	"go.uber.org/zap"
)

var cfg struct {
	Logger struct {
		Level            string   `yaml:"level"`
		Encoding         string   `yaml:"encoding"`
		OutputPaths      []string `yaml:"outputPaths"`
		ErrorOutputPaths []string `yaml:"errorOutputPaths"`
	} `yaml:"logger"`
	// ... other configuration fields
}

var Logger *zap.Logger

func Init() {
	l, _ := zap.NewDevelopment()
	Logger = l
}
