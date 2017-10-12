package BluesoundAPI

import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml"
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

func main(){
	var XMLData string
	XMLData = "'<?xml version=1.0 encoding=UTF-8 standalone=yes?><SyncStatus icon=/images/players/N110_nt.png volume=-1 modelName=NODE 2 name=Cave model=N110 brand=Bluesound etag=3 schemaVersion=15 syncStat=3 id=192.168.1.45:11000 mac=90:56:82:3F:AA:A0></SyncStatus>"
	var MySyncStatus BluesoundSyncStatus
	xml.Unmarshal([]byte(XMLData), &MySyncStatus)
	fmt.Println("Name:" + MySyncStatus.Name)
}