package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"pub-sub-service/model"
	"pub-sub-service/service"
)

type pubSubController struct {
	service service.PubSubService
}

func NewPubSubController(service service.PubSubService) pubSubController {
	return pubSubController{service: service}
}

func (controller pubSubController) Publish(ctx *gin.Context) {
	var request model.Event
	err := ctx.ShouldBindWith(&request, binding.JSON)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	controller.service.AddEvent(request)
	controller.service.Broadcast()

	ctx.JSON(http.StatusOK, "")
}
