package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suguan/template-api/_internal"
)

type NotFoundController struct {
	router *gin.Engine
}

func NewNotFoundController(
	router *gin.Engine,
) *NotFoundController {
	return &NotFoundController{
		router: router,
	}
}

func (msg *NotFoundController) RegisterRoute() {
	msg.router.NoRoute(msg.notFound)
}

func (msg *NotFoundController) notFound(c *gin.Context) {
	c.JSON(404, &_internal.ResponseBody{
		Code:    404,
		Message: "not found",
	})
}
