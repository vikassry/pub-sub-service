package subscriber

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"pub-sub-service/constants"
	"pub-sub-service/model"
	"testing"
	"time"
)

var printCallCount = 0

var tempPrint func(format string, a ...interface{}) (n int, err error)

var mockPrint = func(format string, a ...interface{}) (n int, err error) {
	printCallCount += 1
	return printCallCount, nil
}

type JiraSubscriberTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestJiraSubscriberTestSuite(t *testing.T) {
	suite.Run(t, new(JiraSubscriberTestSuite))
}

func (suite *JiraSubscriberTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	printCallCount = 0
	tempPrint = print
	print = mockPrint
}

func (suite *JiraSubscriberTestSuite) TearDownTest() {
	print = tempPrint
	printCallCount = 0
	suite.mockCtrl.Finish()
}

func (suite *JiraSubscriberTestSuite) TestJiraSubscriberShouldHaveNoEventsInitially() {
	subscriber := JiraSubscriber{}
	var expected []model.Event
	suite.Equal(expected, subscriber.getEvents())
}

func (suite *JiraSubscriberTestSuite) TestJiraSubscriberShouldHaveEventsWhenReceived() {
	subscriber := JiraSubscriber{}
	event := model.Event{
		Type: constants.INCIDENT_CREATED,
		Data: model.EventData{},
	}
	subscriber.AddEvent(event)
	expected := []model.Event{event}
	suite.Equal(expected, subscriber.getEvents())
}

func (suite *JiraSubscriberTestSuite) TestJiraSubscriberShouldProcessNewEventsWhenReceived() {
	subscriber := JiraSubscriber{}

	event := model.Event{
		Type: constants.INCIDENT_CREATED,
		Data: model.EventData{
			Id:          1,
			Message:     "new Jira incident",
			Description: "",
			Status:      "created",
			Time:        time.Now().Format(constants.EVENT_TIME_FORMAT),
		},
	}
	subscriber.AddEvent(event)
	subscriber.Process()

	suite.Equal(1, printCallCount)
}

func (suite *JiraSubscriberTestSuite) TestJiraSubscriberShouldProcessOldEventsForIncidentCreated() {
	subscriber := JiraSubscriber{}
	event := model.Event{
		Type: constants.INCIDENT_CREATED,
		Data: model.EventData{
			Id:          1,
			Message:     "new Jira incident",
			Description: "",
			Status:      "created",
			Time:        "2021-10-18T08:20:55:555",
		},
	}
	subscriber.AddEvent(event)

	subscriber.Process()

	suite.Equal(2, printCallCount)
}
