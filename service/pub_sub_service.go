package service

import (
	"fmt"
	"pub-sub-service/model"
	"pub-sub-service/subscriber"
)

type PubSubService interface {
	AddEvent(event model.Event)
	AddSubscriber(topic string, subscriber subscriber.Subscriber)
	RemoveSubscriber(topic string, subscriber subscriber.Subscriber)
	Broadcast()
}

type pubSubService struct {
	events             []model.Event
	topicSubscriberMap map[string][]subscriber.Subscriber
}

func NewPubSubService() PubSubService {
	return &pubSubService{
		events:             make([]model.Event, 0),
		topicSubscriberMap: make(map[string][]subscriber.Subscriber, 0),
	}
}

func (service *pubSubService) AddEvent(event model.Event) {
	service.events = append(service.events, event)
}

func (service *pubSubService) getEvents() []model.Event {
	return service.events
}

func (service *pubSubService) AddSubscriber(topic string, subscriber subscriber.Subscriber) {
	service.topicSubscriberMap[topic] = append(service.topicSubscriberMap[topic], subscriber)
}

func (service *pubSubService) getSubscribers(topic string) []subscriber.Subscriber {
	return service.topicSubscriberMap[topic]
}

func (service *pubSubService) RemoveSubscriber(topic string, sub subscriber.Subscriber) {
	subscribers := service.topicSubscriberMap[topic]
	removedSubscribers := make([]subscriber.Subscriber, len(subscribers))
	index := 0
	for _, s := range subscribers {
		if s != sub {
			removedSubscribers[index] = s
			index++
		}
	}
	service.topicSubscriberMap[topic] = removedSubscribers[:index]
}

func (service *pubSubService) Broadcast() {
	if len(service.events) == 0 {
		fmt.Println("No events available to broadcast")
	} else {
		for len(service.events) != 0 {
			event := service.events[0]
			service.events = service.events[1:]
			subscribers, ok := service.topicSubscriberMap[event.Type]
			if !ok {
				fmt.Printf("No subscribers available for event %s", event.Type)
			} else {
				fmt.Printf("Found %d subscribers for event %s", len(subscribers), event.Type)
				for _, sub := range subscribers {
					sub.AddEvent(event)
					sub.Process()
				}
			}
		}
	}
}
