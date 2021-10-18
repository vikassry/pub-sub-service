package subscriber

import (
	"fmt"
	"pub-sub-service/constants"
	"pub-sub-service/model"
	"time"
)

type JiraSubscriber struct {
	events []model.Event
}

const fiveMin = 5

var print = fmt.Printf

func (jira *JiraSubscriber) Process() {
	for len(jira.events) != 0 {
		event := jira.events[0]
		print("\nprocessing event for JIRA %+v\n", event)
		jira.events = jira.events[1:]
		if event.Type == constants.INCIDENT_CREATED {
			eventCreationTime, _ := time.Parse(constants.EVENT_TIME_FORMAT, event.Data.Time)
			now := time.Now()

			if now.Sub(eventCreationTime).Minutes() > float64(fiveMin) {
				print("A JIRA ticket %s has been successfully created\n", event.Data.Message)
			}
		}
	}
}

func (jira *JiraSubscriber) AddEvent(event model.Event) {
	jira.events = append(jira.events, event)
}

func (jira *JiraSubscriber) getEvents() []model.Event {
	return jira.events
}
