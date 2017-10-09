package main

import (
	"log"
	"net/http"

	"google.golang.org/api/calendar/v3"
)

type Calendar struct {
	service *calendar.Service
}

func NewCalendar(c *http.Client) *Calendar {
	s, err := calendar.New(c)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return &Calendar{service: s}
}

func (c *Calendar) Events() ([]*calendar.Event, error) {
	evts, err := c.service.Events.List("primary").Do()
	if err != nil {
		log.Fatalf("%v", err)
	}
	return evts.Items, nil
}
