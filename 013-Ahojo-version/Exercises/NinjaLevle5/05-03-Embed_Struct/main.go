package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	trucker := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Purple",
		},
		fourWheel: true,
	}

	richGuy := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(trucker.doors, trucker.color, trucker.fourWheel)
	fmt.Println(richGuy.doors, richGuy.color, richGuy.luxury)

	// Print out a single field

}
