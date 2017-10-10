package main

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

type Formatter struct{}

func (f *Formatter) Event(e *calendar.Event) string {
	var d string
	if e.Start.DateTime != "" {
		d = e.Start.DateTime
	} else {
		d = e.Start.Date
	}
	return fmt.Sprintf("%s (%s)", e.Summary, d)
}
