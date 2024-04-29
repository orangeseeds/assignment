package main

import (
	"fmt"

	"github.com/orangeseeds/berrybytes/pkg"
)

func main() {
	doorBell := pkg.NewDoorBell()
	doorBell.SetVolume(100)
	doorBell.OnBellRing().Add(func(e *pkg.BellRingEvent) error {
		// Let's say  this is a trigger to activate the door's dash-cam to start recording when the doorbell rings.
		fmt.Println("Door's dash-cam started recording...")
		return nil
	})
	chime, vol, err := doorBell.Ring()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("playing chime: %s at volume %d\n", chime, vol)
}
