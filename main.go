package main

import (
	"fmt"
	"log"
)

func main() {
	c := &Calendar{Client: &Client{}}
	evts, err := c.Events()
	if err != nil {
		log.Fatalf("%v", err)
	}

	if len(evts) > 0 {
		f := &Formatter{}
		for _, e := range evts {
			fmt.Println(f.Format(e))
		}
	} else {
		fmt.Printf("No events found.\n")
	}
}
