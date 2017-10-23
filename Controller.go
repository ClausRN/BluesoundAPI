package BluesoundAPI

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// NewBluesoundController is creating a new controller
func NewBluesoundController(FQDN string) (blueControl BluesoundController) {
	blueControl.baseURL = "http://" + FQDN + ":11000/"

	blueControl.stopSignal = make(chan os.Signal, 1)
	signal.Notify(blueControl.stopSignal, os.Interrupt, os.Kill)
	defer signal.Stop(blueControl.stopSignal)

	return blueControl
}

// Start stops the internal updates
func (blueControl *BluesoundController) Start() bool {
	log.Println("About to start")
	go blueControl.updateData()
	for blueControl.pause == 0 {
		time.Sleep(50 * time.Millisecond)
	}
	log.Println("Started")
	return true
}

// Close stops the internal updates
func (blueControl *BluesoundController) Close() bool {
	log.Println("About to stop")
	close(blueControl.stopSignal)
	log.Println("Thread signaled")
	return true
}

// Status returns current status of player
func (blueControl *BluesoundController) Status() BluesoundStatus {
	statusUpdate.RLock()
	temp := blueControl.status
	statusUpdate.RUnlock()
	return temp
}

// SyncStatus returns current status of player
func (blueControl *BluesoundController) SyncStatus() BluesoundSyncStatus {
	statusUpdate.RLock()
	temp := blueControl.syncStatus
	statusUpdate.RUnlock()
	return temp
}

// GetVersion returns the players software version
func (blueControl *BluesoundController) GetVersion() (Version string) {
	playstate := BluesoundPlayerVersion{}
	if XMLDataBin, err := blueControl.getContent(httpURIVersion); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		} else {
			Version = playstate.Version
		}
	}
	return
}

// Play starts the player from current track
func (blueControl *BluesoundController) Play() (State BluesoundCommandState) {
	return blueControl.sendCommandPlayPause(httpURIPlay)
}

// Pause the player at current position
func (blueControl *BluesoundController) Pause() (State BluesoundCommandState) {
	return blueControl.sendCommandPlayPause(httpURIPause)
}

// Skip to next track
func (blueControl *BluesoundController) Skip() (State BluesoundCommandStateSkipBack) {
	return blueControl.sendCommandSkipBack(httpURISkip)
}

// Back to start or previous track
func (blueControl *BluesoundController) Back() (State BluesoundCommandStateSkipBack) {
	return blueControl.sendCommandSkipBack(httpURIBack)
}

// sendCommandPlayPause sends a simple command to the player
func (blueControl *BluesoundController) sendCommandPlayPause(Command string) (State BluesoundCommandState) {
	playstate := BluesoundCommandState{}
	playstate.State = "No State"
	if XMLDataBin, err := blueControl.getContent(Command); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		}
	}
	log.Printf("Player command: %s - state: %s", Command, playstate.State)
	return playstate
}

// sendCommandSkipBack sends a simple command to the player
func (blueControl *BluesoundController) sendCommandSkipBack(Command string) (State BluesoundCommandStateSkipBack) {
	playstate := BluesoundCommandStateSkipBack{}
	if XMLDataBin, err := blueControl.getContent(Command); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		}
	}
	log.Printf("Player command: %s - ID: %d", Command, playstate.TrackID)
	return playstate
}

// Clear the playlist
func (blueControl *BluesoundController) Clear() (State bool) {
	var success = false
	playstate := BluesoundPlayQueue{}
	if XMLDataBin, err := blueControl.getContent(httpURIClear); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		} else {
			if playstate.Length == 0 {
				success = true
			}
		}
	}
	return success
}

// SetVolume sets the player output level (0 - 100, 0 = Mute)
func (blueControl *BluesoundController) SetVolume(Level int) (State bool) {
	var success = false
	if Level >= 0 && Level <= 100 {
		playstate := BluesoundVolume{}
		if XMLDataBin, err := blueControl.getContent(httpURISetVolume + "?level=" + strconv.Itoa(Level)); err != nil {
			log.Printf("Failed to get XML: %v", err)
		} else {
			// log.Printf("Received XML answer: %s", string(XMLDataBin))
			err = xml.Unmarshal(XMLDataBin, &playstate)
			if err != nil {
				log.Printf("XML data not valid: %s", err)
			} else {
				if playstate.Volume == Level {
					success = true
				}
			}
		}
	}
	return success
}

// SetShuffle sets the player shuffle mode
func (blueControl *BluesoundController) SetShuffle(Level int) (State bool) {
	var success = false
	if Level >= 0 && Level <= 1 {
		playstate := BluesoundPlayQueue{}
		if XMLDataBin, err := blueControl.getContent(httpURIShuffle + "?state=" + strconv.Itoa(Level)); err != nil {
			log.Printf("Failed to get XML: %v", err)
		} else {
			// log.Printf("Received XML answer: %s", string(XMLDataBin))
			err = xml.Unmarshal(XMLDataBin, &playstate)
			if err != nil {
				log.Printf("XML data not valid: %s", err)
			} else {
				if playstate.Shuffle == Level {
					success = true
				}
			}
		}
	}
	return success
}

// SetRepeat sets the player shuffle mode
func (blueControl *BluesoundController) SetRepeat(Level int) (State bool) {
	var success = false
	if Level >= 0 && Level <= 2 {
		playstate := BluesoundPlayQueue{}
		if XMLDataBin, err := blueControl.getContent(httpURIRepeat + "?state=" + strconv.Itoa(Level)); err != nil {
			log.Printf("Failed to get XML: %v", err)
		} else {
			// log.Printf("Received XML answer: %s", string(XMLDataBin))
			err = xml.Unmarshal(XMLDataBin, &playstate)
			if err != nil {
				log.Printf("XML data not valid: %s", err)
			} else {
				if playstate.Repeat == Level {
					success = true
				}
			}
		}
	}
	return success
}

// GetPlaylists get playlist from player
func (blueControl *BluesoundController) GetPlaylists() (State BluesoundPlaylists) {
	playstate := BluesoundPlaylists{}
	if XMLDataBin, err := blueControl.getContent(httpURIPlaylists); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		}
	}
	return playstate
}

// GetPlaylist get playlist from player
func (blueControl *BluesoundController) GetPlaylist(Playlist string) (State BluesoundPlaylist) {
	playstate := BluesoundPlaylist{}
	if XMLDataBin, err := blueControl.getContent(httpURISongs + "?playlist=" + url.QueryEscape(Playlist)); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		}
	}
	return playstate
}

// GetPlayQueue get playlist from player
func (blueControl *BluesoundController) GetPlayQueue() (State BluesoundPlayQueue) {
	playstate := BluesoundPlayQueue{}
	if XMLDataBin, err := blueControl.getContent(httpURIPlaylist); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		}
	}
	return playstate
}

// PlayPlaylist starts playing the playlist
func (blueControl *BluesoundController) PlayPlaylist(Playlist string) (State BluesoundAddSong) {
	playstate := BluesoundAddSong{}
	if XMLDataBin, err := blueControl.getContent(httpURIAdd + "?playlist=" + url.QueryEscape(Playlist) + "&playnow=-1"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		// log.Printf("Received XML answer: %s", string(XMLDataBin))
		err = xml.Unmarshal(XMLDataBin, &playstate)
		if err != nil {
			log.Printf("XML data not valid: %s", err)
		}
	}
	return playstate
}

// getContent return data from http request
func (blueControl *BluesoundController) getContent(CommandURL string) ([]byte, error) {
	resp, err := http.Get(blueControl.baseURL + CommandURL)
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

// updateData updates the values for status and syncstatus
func (blueControl *BluesoundController) updateData() {
	var Running = true
	var StatusType int
	var StatusURL string
	//var pause time.Duration
	for Running {
		log.Println("Updating data")
		switch StatusType {
		case 0:
			StatusURL = httpURIStatus
		case 1:
			StatusURL = httpURISyncStatus
		}
		XMLDataBin, err := blueControl.getContent(StatusURL)
		if err != nil {
			log.Printf("Could not get %s: %s", StatusURL, err)
		} else {
			statusUpdate.Lock()
			switch StatusType {
			case 0:
				err = xml.Unmarshal(XMLDataBin, &blueControl.status)
			case 1:
				err = xml.Unmarshal(XMLDataBin, &blueControl.syncStatus)
			}
			statusUpdate.Unlock()
			if err != nil {
				log.Printf("XML data not valid for %s: %s", StatusURL, err)
			}
		}
		StatusType++
		if StatusType > 1 {
			StatusType = 0
			blueControl.pause = 500
		}
		select {
		case <-time.After(blueControl.pause * time.Millisecond):
			if 1 == 0 {
				log.Println("Dummy")
			}
		case <-blueControl.stopSignal:
			log.Println("Stop signal received, exiting!")
			Running = false
		}
	}
	log.Println("Stopped")
}
