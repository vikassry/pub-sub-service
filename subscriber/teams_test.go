package subscriber

import (
	"github.com/stretchr/testify/assert"
	"pub-sub-service/constants"
	"pub-sub-service/model"
	"testing"
)

func TestTeamsSubscriberShouldHaveNoEventsInitially(t *testing.T) {
	subscriber := TeamsSubscriber{}
	var expected []model.Event
	assert.Equal(t, expected, subscriber.getEvents())
}

func TestTeamsSubscriberShouldHaveEventsWhenReceived(t *testing.T) {
	subscriber := TeamsSubscriber{}
	event := model.Event{
		Type: constants.INCIDENT_ACKNOWLEDGED,
		Data: model.EventData{},
	}
	subscriber.AddEvent(event)
	expected := []model.Event{event}
	assert.Equal(t, expected, subscriber.getEvents())
}
