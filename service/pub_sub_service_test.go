package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"pub-sub-service/constants"
	"pub-sub-service/mocks"
	"pub-sub-service/model"
	"pub-sub-service/subscriber"
	"testing"
)

func TestPubSubServiceShouldAddEventsSuccessfully(t *testing.T) {
	pubSubService := pubSubService{}
	pubSubService.AddEvent(model.Event{Type: constants.INCIDENT_CREATED})
	pubSubService.AddEvent(model.Event{Type: constants.INCIDENT_ACKNOWLEDGED})
	expected := []model.Event{{Type: constants.INCIDENT_CREATED}, {Type: constants.INCIDENT_ACKNOWLEDGED}}

	assert.Equal(t, expected, pubSubService.getEvents())
}

func TestPubSubServiceShouldAddSubscriberSuccessfully(t *testing.T) {
	service := NewPubSubService()
	service.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, &subscriber.SlackSubscriber{})
	service.AddSubscriber(constants.INCIDENT_CREATED, &subscriber.TeamsSubscriber{})
	service.AddSubscriber(constants.INCIDENT_CREATED, &subscriber.TeamsSubscriber{})

	slack := service.(*pubSubService).getSubscribers(constants.INCIDENT_ACKNOWLEDGED)
	teams := service.(*pubSubService).getSubscribers(constants.INCIDENT_CREATED)

	assert.Equal(t, []subscriber.Subscriber{&subscriber.SlackSubscriber{}}, slack)
	assert.Equal(t, []subscriber.Subscriber{&subscriber.TeamsSubscriber{}, &subscriber.TeamsSubscriber{}}, teams)
}

func TestPubSubServiceShouldRemoveSubscriberSuccessfully(t *testing.T) {
	service := NewPubSubService()
	subToBeRemoved := &subscriber.SlackSubscriber{}
	service.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, &subscriber.SlackSubscriber{})
	service.AddSubscriber(constants.INCIDENT_CREATED, &subscriber.TeamsSubscriber{})
	service.AddSubscriber(constants.INCIDENT_CREATED, subToBeRemoved)

	service.RemoveSubscriber(constants.INCIDENT_CREATED, subToBeRemoved)

	slack := service.(*pubSubService).getSubscribers(constants.INCIDENT_ACKNOWLEDGED)
	teams := service.(*pubSubService).getSubscribers(constants.INCIDENT_CREATED)

	assert.Equal(t, []subscriber.Subscriber{&subscriber.SlackSubscriber{}}, slack)
	assert.Equal(t, []subscriber.Subscriber{&subscriber.TeamsSubscriber{}}, teams)
}

func TestPubSubServiceShouldBroadcastEventToSlackSubscriber(t *testing.T) {
	service := NewPubSubService()
	mockSlack := mocks.NewMockSubscriber(gomock.NewController(t))
	service.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, mockSlack)
	service.AddSubscriber(constants.INCIDENT_CREATED, &subscriber.TeamsSubscriber{})
	service.AddEvent(model.Event{Type: constants.INCIDENT_ACKNOWLEDGED})
	service.AddEvent(model.Event{Type: constants.INCIDENT_ACKNOWLEDGED})

	mockSlack.EXPECT().AddEvent(model.Event{Type: constants.INCIDENT_ACKNOWLEDGED}).Times(2)
	mockSlack.EXPECT().Process().Times(2)

	service.Broadcast()

	events := service.(*pubSubService).getEvents()
	assert.Equal(t, []model.Event{}, events)
}

func TestPubSubServiceShouldNotBroadcastWhenNoEventsPublished(t *testing.T) {
	service := NewPubSubService()
	mockSlack := mocks.NewMockSubscriber(gomock.NewController(t))
	service.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, mockSlack)

	mockSlack.EXPECT().AddEvent(model.Event{Type: constants.INCIDENT_ACKNOWLEDGED}).Times(0)
	mockSlack.EXPECT().Process().Times(0)
	service.Broadcast()
}

func TestPubSubServiceShouldNotBroadcastWhenNoSubscribersAvailableForATopic(t *testing.T) {
	mockSlack := mocks.NewMockSubscriber(gomock.NewController(t))
	service := NewPubSubService()

	service.AddSubscriber(constants.INCIDENT_ACKNOWLEDGED, mockSlack)
	service.AddEvent(model.Event{Type: constants.INCIDENT_CREATED})

	mockSlack.EXPECT().AddEvent(model.Event{Type: constants.INCIDENT_ACKNOWLEDGED}).Times(0)
	mockSlack.EXPECT().Process().Times(0)

	service.Broadcast()
}
