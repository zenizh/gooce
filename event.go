package main

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

type Event struct {
	summary  string
	date     string
	dateTime string
}

func NewEvent(event *calendar.Event) *Event {
	return &Event{
		summary:  event.Summary,
		date:     event.Start.Date,
		dateTime: event.Start.DateTime,
	}
}

func (event *Event) String() string {
	var date string
	if event.dateTime != "" {
		date = event.dateTime
	} else {
		date = event.date
	}
	return fmt.Sprintf("%s (%s)", event.summary, date)
}
