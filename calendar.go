package main

import (
	"net/http"
	"time"

	"google.golang.org/api/calendar/v3"
)

type Calendar struct {
	EventList []*Event

	service *calendar.Service
}

func NewCalendar(client *http.Client) (*Calendar, error) {
	service, err := calendar.New(client)
	if err != nil {
		return nil, err
	}
	return &Calendar{service: service}, nil
}

func (calendar *Calendar) FetchEvents() error {
	t := time.Now().Format(time.RFC3339)

	evts, err := calendar.service.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("StartTime").Do()
	if err != nil {
		return err
	}

	for _, item := range evts.Items {
		calendar.EventList = append(calendar.EventList, NewEvent(item))
	}

	return nil
}
