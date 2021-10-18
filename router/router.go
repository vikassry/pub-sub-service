package router

import (
	"github.com/gin-gonic/gin"
	"pub-sub-service/constants"
	"pub-sub-service/controller"
	"pub-sub-service/service"
	"pub-sub-service/subscriber"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	group := router.Group("/api/pub-sub")
	{
		slackSubscriber := subscriber.SlackSubscriber{}
		teamsSubscriber := subscriber.TeamsSubscriber{}
		jiraSubscriber := subscriber.JiraSubscriber{}
		pubSubService := service.NewPubSubService()
		pubSubService.AddSubscriber(constants.INCIDENT_CREATED, &slackSubscriber)
		pubSubService.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, &slackSubscriber)

		pubSubService.AddSubscriber(constants.INCIDENT_RESOLVED, &teamsSubscriber)
		pubSubService.AddSubscriber(constants.INCIDENT_REASSIGNED, &teamsSubscriber)

		pubSubService.AddSubscriber(constants.INCIDENT_CREATED, &jiraSubscriber)
		pubSubService.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, &jiraSubscriber)
		pubSubService.AddSubscriber(constants.INCIDENT_RESOLVED, &jiraSubscriber)

		pubSubController := controller.NewPubSubController(pubSubService)

		group.POST("/publish", pubSubController.Publish)
	}
	return router
}
