package BluesoundAPI

import (
	"log"
)

// NewBluesoundController is creating a new controller
func NewBluesoundController() (blueControl BluesoundController) {
	blueControl.Name = ""
	return blueControl
}

// Play starts the player from current track
func (blueControl *BluesoundController) Play() bool {
	blueControl.Name = "Play"
	log.Println("Playing")
	return true
}
