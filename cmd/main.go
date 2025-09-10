package main

import (
	"go_mail_sender/connectors"
	"go_mail_sender/controller"
	"go_mail_sender/repository"
	"log"

	"os"

	"github.com/labstack/echo"
)

type Container struct {
	MailInstance controller.MailController
}

func LoadContainer() *Container {
	return &Container{
		MailInstance: controller.MailController{Mysql: repository.SingleTonPattern()},
	}
}

func init() {
	connectors.LoadEnv()
}

func main() {
	e := echo.New()
	PORT := os.Getenv("PORT")

	containerInstance := LoadContainer()
	e.POST("/send", containerInstance.MailInstance.SendMail)
	e.GET("/logs/mysql", containerInstance.MailInstance.GetMailLogs)

	if PORT == "" {
		PORT = "9090"

	}
	log.Fatal(e.Start(":" + PORT))
}
