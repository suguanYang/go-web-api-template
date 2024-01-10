package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"github.com/suguan/wechat-hm/controller"
	"github.com/suguan/wechat-hm/db"
	"github.com/suguan/wechat-hm/repository"
)

func main() {
	container := BuildContainer()
	err := container.Invoke(func(router *gin.Engine, controller *controller.MessageController, notFoundController *controller.NotFoundController) {
		router.Use(gin.Logger())
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		controller.RegisterRoute()
		notFoundController.RegisterRoute()

		router.Run(":8080")
	})

	if err != nil {
		panic(err)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(db.Setup)
	_ = container.Provide(gin.Default)
	_ = container.Provide(repository.NewMessageRepository)
	_ = container.Provide(controller.NewMessageController)
	_ = container.Provide(controller.NewNotFoundController)
	return container
}
