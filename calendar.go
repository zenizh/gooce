package main

type Calendar struct {
	Client *Client
}

func (c *Calendar) Events() ([]Event, error) {
	return make([]Event, 10), nil
}
