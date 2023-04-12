package main

import (
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/seaung/gin-machinery/server"
	"github.com/seaung/gin-machinery/utils"
	"github.com/seaung/gin-machinery/worker"
	"github.com/urfave/cli"
)

var (
	app             *cli.App
	machineryServer *machinery.Server
)

func init() {
	app = cli.NewApp()
	app.Name = "gin-machinery"
	app.Author = "seaung"

	machineryServer = utils.NewMachineryServer()
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "开启web服务并运行machinery后台任务",
			Action: func(c *cli.Context) {
				utils.Logger.Info("running server ...")
				server.RunServer(machineryServer)
			},
		},
		{
			Name:  "worker",
			Usage: "运行工作者服务",
			Action: func(c *cli.Context) {
				utils.Logger.Info("worker ...")
				worker.RunWorker(machineryServer)
			},
		},
	}

	app.Run(os.Args)
}
