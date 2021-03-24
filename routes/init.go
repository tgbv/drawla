package routes

import (
	"github.com/tgbv/drawla/controllers"
	"goji.io"
	"goji.io/pat"
)

/*
*	default function which receives the goji mux and registers all actuve routes with controllers
 */
func Init(m *goji.Mux) {

	// homepage
	m.HandleFunc(pat.Get("/"), controllers.Home)

	// create a new room
	m.HandleFunc(pat.Get("/room/new"), controllers.NewRoom)

	// join custom room
	m.HandleFunc(pat.Get("/room/:room"), controllers.JoinRoom)

	//

}
