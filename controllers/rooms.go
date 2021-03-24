package controllers

import (
	"fmt"
	"net/http"

	"goji.io/pat"

	"github.com/dchest/uniuri"
	"github.com/tgbv/drawla/lib"
	"github.com/tgbv/drawla/state"
	s "github.com/tgbv/drawla/state"
)

/*
*	joins to an existing room
 */
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	room := pat.Param(r, "room")

	// check if room exists
	if state.Rooms[room] == nil {
		fmt.Fprintf(w, "Room %s does not exist!", room)
		return
	}

	t := lib.GetTmpl("draw.htm")
	t.Execute(w, map[string]interface{}{
		"room_id": room,
	})
}

/*
*	creates a new room
 */
func NewRoom(w http.ResponseWriter, r *http.Request) {
	str := uniuri.NewLen(10)

	s.Rooms.Add(str)
	fmt.Println("New room:", str)

	// t := lib.GetTmpl("draw.htm")
	// t.Execute(w, map[string]string{
	// 	"room_id": str,
	// })

	w.Header().Add("Location", "/room/"+str)
	w.WriteHeader(302)
}
