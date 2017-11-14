package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func NewClient() (*http.Client, error) {
	json, err := ioutil.ReadFile(filepath.Join(baseDir(), "client_secret.json"))
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(json, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	token, err := restoreToken()
	if err != nil {
		token, err = requestToken(config)
		if err != nil {
			return nil, err
		}

		err = saveToken(token)
		if err != nil {
			return nil, err
		}
	}

	return config.Client(context.Background(), token), nil
}

func restoreToken() (*oauth2.Token, error) {
	file, err := os.Open(filepath.Join(baseDir(), "token.json"))
	if err != nil {
		return nil, err
	}

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	defer file.Close()

	return token, err
}

func requestToken(config *oauth2.Config) (*oauth2.Token, error) {
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf(`1. Go to the following link in your browser:
%v

2. Type the authorization code: `, url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func saveToken(token *oauth2.Token) error {
	file, err := os.OpenFile(filepath.Join(baseDir(), "token.json"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	json.NewEncoder(file).Encode(token)
	return nil
}

func baseDir() string {
	var dir string
	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("APPDATA")
	default:
		dir = os.Getenv("HOME")
	}
	return filepath.Join(dir, ".gooce")
}
