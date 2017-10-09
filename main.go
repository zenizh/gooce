package main

import (
	"fmt"
	"log"
)

func main() {
	c := NewCalendar(NewClient())

	evts, err := c.Events()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if len(evts) > 0 {
		for _, e := range evts {
			fmt.Println(e)
		}
	} else {
		fmt.Printf("No events found.\n")
	}
}
