package main

import (
	"log"
	"net/http"
	"time"

	"google.golang.org/api/calendar/v3"
)

type Calendar struct {
	service *calendar.Service
}

func NewCalendar(c *http.Client) *Calendar {
	s, err := calendar.New(c)
	if err != nil {
		log.Fatal(err)
	}
	return &Calendar{service: s}
}

func (c *Calendar) Events() ([]*calendar.Event, error) {
	t := time.Now().Format(time.RFC3339)
	evts, err := c.service.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("StartTime").Do()
	if err != nil {
		log.Fatal(err)
	}
	return evts.Items, nil
}
