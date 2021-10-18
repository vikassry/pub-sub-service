package subscriber

import (
	"pub-sub-service/model"
)

type Subscriber interface {
	AddEvent(event model.Event)
	Process()
}
