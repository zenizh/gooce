package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const (
	baseDir    = ".goce"
	secretFile = "client_secret.json"
	tokenFile  = "token.json"
)

func NewClient() *http.Client {
	b, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	c, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("%v", err)
	}

	f, err := tokenFilePath()
	if err != nil {
		log.Fatalf("%v", err)
	}

	t, err := tokenFromFile(f)
	if err != nil {
		t = requestToken(c)
		saveToken(f, t)
	}

	return c.Client(context.Background(), t)
}

func tokenFilePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	d := filepath.Join(u.HomeDir, baseDir)
	os.MkdirAll(d, 0700)
	return filepath.Join(d, url.QueryEscape(tokenFile)), nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()

	return t, nil
}

func requestToken(c *oauth2.Config) *oauth2.Token {
	URL := c.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf(`1. Go to the following link in your browser:
%v

2. Type the authorization code: `, URL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("%v", err)
	}

	t, err := c.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return t
}

func saveToken(file string, t *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(t)
}
