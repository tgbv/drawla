package controllers

import (
	"fmt"
	"goji/pat"
	"net/http"
)

/*
*	joins to a new or existing room
 */
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	room := pat.Param(r)

	fmt.Fprintf(w, "Room: %g", room)
}
