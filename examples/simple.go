package main

import (
	"fmt"

	"github.com/rcmendes/go-logs/pkg/shared/logs"
)

//TODO Create log level and init function

func main() {
	builder := logs.NewLoggerBuilder()

	logger := builder.Build()
	fmt.Println("Logging with default level:")
	generateLogs(logger)

	builder.SetLevel(logs.LevelDebug)
	logger = builder.Build()
	fmt.Println("\nLogging in DEBUG level:")
	generateLogs(logger)

	builder.SetLevel(logs.LevelInfo)
	logger = builder.Build()
	fmt.Println("\nLogging in INFO level:")
	generateLogs(logger)

	builder.SetLevel(logs.LevelWarning)
	logger = builder.Build()
	fmt.Println("\nLogging in WARNING level:")
	generateLogs(logger)

	builder.SetLevel(logs.LevelError)
	logger = builder.Build()
	fmt.Println("\nLogging in ERROR level:")
	generateLogs(logger)

}

func generateLogs(logger logs.Logger) {
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warning("Warning message")
	logger.Error("Error message")
}
