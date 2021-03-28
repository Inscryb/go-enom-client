package enom

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	Session *Session
}

func (c Client) DoRequest(command *Command) ([]byte, error) {
	doURL, err := url.Parse(c.Session.ResellerURL + "/interface.asp")
	if err != nil {
		log.Fatal(err)
	}
	doURL.RawQuery = command.ToParams()

	urlString := doURL.String()
	log.Print("URL to be called: ", urlString)

	resp, err := http.Get(urlString)
	if err != nil {
		log.Fatal("Post Failed: ", err)
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read: ", err)
		return nil, err
	}

	//log.Print(resp.Status)
	//log.Print(fmt.Printf("%s\n", body))

	return body, nil
}
