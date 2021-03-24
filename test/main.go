package main

import (
	"fmt"

	l "github.com/tgbv/drawla/lib"
)

func main() {

	// tests rooms with points
	rooms := make(l.Rooms)
	rn := "alehandro"

	rooms.Add(rn)
	rooms[rn].AddPoint(22, 33)
	rooms[rn].DeletePoint(func(p l.Point) bool {
		if p.Coords.X == 232 && p.Coords.Y == 33 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(rooms[rn].Points)
}
