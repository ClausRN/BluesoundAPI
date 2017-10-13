package BluesoundAPI

import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
//	"encoding/xml"
)

// NewBluesoundController is creating a new controller
func NewBluesoundController(FQDN string) (blueControl BluesoundController) {
	blueControl.Name = ""
	blueControl.baseURL = "http://" + FQDN + ":11000/"

	if data, err := blueControl.getContent(blueControl.baseURL + "Status"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println("Received XML:")
		log.Println(string(data))
	}
	
	return blueControl
}

// Play starts the player from current track
func (blueControl *BluesoundController) Play() bool {
	blueControl.Name = "Play"
	log.Println("Playing")
	return true
}

// getContent return data from http request
func (blueControl *BluesoundController) getContent(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Read body: %v", err)
    }

    return data, nil
}

