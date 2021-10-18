package subscriber

import (
	"fmt"
	"pub-sub-service/model"
)

type SlackSubscriber struct {
	events []model.Event
}

func (slack *SlackSubscriber) Process() {
	for len(slack.events) != 0 {
		event := slack.events[0]
		slack.events = slack.events[1:]
		fmt.Printf("\nprocessing event for Slack %+v\n", event)
	}
}

func (slack *SlackSubscriber) AddEvent(event model.Event) {
	slack.events = append(slack.events, event)
}

func (slack *SlackSubscriber) getEvents() []model.Event {
	return slack.events
}
