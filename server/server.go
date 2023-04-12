package server

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
	task "github.com/seaung/gin-machinery/tasks"
)

func RunServer(taskserver *machinery.Server) {
	app := gin.New()

	app.POST("/send_message", func(c *gin.Context) {
		parameters := new(task.Message)
		if err := c.ShouldBindJSON(parameters); err != nil {
			log.Println(err)
		}

		req, err := json.Marshal(parameters)
		if err != nil {
			log.Println(err)
		}

		str := base64.StdEncoding.EncodeToString([]byte(req))
		tk := tasks.Signature{
			Name: "send_message",
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: str,
				},
			},
		}

		resp, err := taskserver.SendTask(&tk)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"task_id":   resp.GetState().TaskUUID,
			"task_name": resp.GetState().TaskName,
		})
	})

	app.Run(":9000")
}
