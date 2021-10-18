package subscriber

import (
	"fmt"
	"pub-sub-service/model"
)

type TeamsSubscriber struct {
	events []model.Event
}

func (teams *TeamsSubscriber) Process() {
	for len(teams.events) != 0 {
		event := teams.events[0]
		teams.events = teams.events[1:]
		fmt.Printf("\nprocessing event for Teams %+v\n", event)
	}
}

func (teams *TeamsSubscriber) AddEvent(event model.Event) {
	teams.events = append(teams.events, event)
}

func (teams *TeamsSubscriber) getEvents() []model.Event {
	return teams.events
}
