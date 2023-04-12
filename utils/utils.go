package utils

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	task "github.com/seaung/gin-machinery/tasks"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
}

func NewMachineryServer() *machinery.Server {
	Logger.Info("new machinery instance")

	taskserver, err := machinery.NewServer(&config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	})

	if err != nil {
		Logger.Fatal(err.Error())
	}

	taskserver.RegisterTasks(map[string]interface{}{
		"send_message": task.SendMessage,
	})

	return taskserver
}
