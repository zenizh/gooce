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
	c, err := google.ConfigFromJSON(readSecretFile(), calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatal(err)
	}

	f := tokenFilePath()

	t, err := tokenFromFile(f)
	if err != nil {
		t = requestToken(c)
		saveToken(f, t)
	}

	return c.Client(context.Background(), t)
}

func readSecretFile() []byte {
	b, err := ioutil.ReadFile(filepath.Join(baseDirPath(), secretFile))
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func baseDirPath() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	d := filepath.Join(u.HomeDir, baseDir)
	os.MkdirAll(d, 0700)
	return d
}

func tokenFilePath() string {
	return filepath.Join(baseDirPath(), url.QueryEscape(tokenFile))
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func requestToken(c *oauth2.Config) *oauth2.Token {
	URL := c.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf(`1. Go to the following link in your browser:
%v

2. Type the authorization code: `, URL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	t, err := c.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func saveToken(file string, t *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(t)
}
