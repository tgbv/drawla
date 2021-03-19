package controllers

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

/*
*	joins to a new or existing room
 */
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	room := pat.Param(r, "room")

	fmt.Fprintf(w, "Room: %s", room)
}

/*
*	creates a new room
 */
func NewRoom(w http.ResponseWriter, r *http.Request) {

}
