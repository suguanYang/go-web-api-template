package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suguan/template-api/_internal"
	"github.com/suguan/template-api/repository"
)

type MessageController struct {
	repo      repository.MessageRepositoryI
	router    *gin.Engine
	baseRoute string
}

func NewMessageController(
	repo repository.MessageRepositoryI,
	router *gin.Engine,
) *MessageController {
	return &MessageController{
		repo:      repo,
		router:    router,
		baseRoute: "/api",
	}
}

func (msg *MessageController) RegisterRoute() {
	routeV1Group := msg.router.Group(msg.baseRoute + "/v1")

	routeV1Group.GET("/talker/:talker_id/message", msg.getTalkerMessage)
}

func (msg *MessageController) getTalkerMessage(c *gin.Context) {
	talkerID := c.Param("talker_id")
	pagination := _internal.PaginationQuery{}
	err := c.ShouldBindQuery(&pagination)
	if err != nil {
		panic(err)
	}

	messages := msg.repo.GetMessageListByTalkerID(talkerID, pagination)

	c.JSON(200, &_internal.ListResponseBody{
		ResponseBody: _internal.ResponseBody{
			Code:    0,
			Message: "success",
			Data:    messages,
		},
		Pagination: &_internal.PaginationResponse{
			Total: len(messages),
			Page:  pagination.Page,
		},
	})
}
