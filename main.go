package main

import (
	"fmt"

	"github.com/orangeseeds/berrybytes/pkg"
)

func main() {
	doorBell := pkg.NewDoorBell()
	doorBell.SetVolume(100)
	doorBell.OnBellRing().Add(func(e *pkg.BellRingEvent) error {
		fmt.Printf("Bell rang with '%v' sound, running some other action!!\n", e.Sound)
		return nil
	})
	fmt.Println("ring ring ring, ", doorBell.Ring())
}
