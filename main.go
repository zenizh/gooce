package main

import (
	"fmt"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	calendar, err := NewCalendar(client)
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	err = calendar.FetchEvents()
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	if len(calendar.EventList) > 0 {
		for _, e := range calendar.EventList {
			fmt.Println(e)
		}
	} else {
		fmt.Println("No events found.")
	}

	return ExitCodeOK
}
