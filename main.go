package main

import (
	"fmt"
	"net/http"

	"github.com/ambelovsky/gosf"
	"github.com/tgbv/drawla/routes"
	"github.com/tgbv/drawla/ws"
	"goji.io"
)

func echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage(request.Message.Text)
}

func moveMouse(c *gosf.Client, r *gosf.Request) *gosf.Message {
	m := new(gosf.Message)
	m.Success = true
	m.Text = "aaaaaaaaaaa"

	//fmt.Println(r.Message)

	gosf.Broadcast("", "mouseMoved", r.Message)

	return m

}

func main() {
	// first init goji mux
	mux := goji.NewMux()

	// pass the mux to the router as reference
	// the routes will be bounded with controllers
	routes.Init(mux)

	// listen HTTP server
	go (func() {
		error := http.ListenAndServe("localhost:8000", mux)
		if error != nil {
			panic(error)
		}
	})()

	// listen ws server
	go ws.Init()

	var input string
	fmt.Println("Press enter to exit program.")
	fmt.Scanln(&input)
}
