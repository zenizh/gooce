package main

import (
	"fmt"
	"log"
)

func main() {
	c := NewCalendar(NewClient())

	evts, err := c.Events()
	if err != nil {
		log.Fatal(err)
	}

	if len(evts) > 0 {
		f := &Formatter{}
		for _, e := range evts {
			fmt.Println(f.Event(e))
		}
	} else {
		fmt.Println("No events found.")
	}
}
