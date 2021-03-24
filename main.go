package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/tgbv/drawla/routes"
	"github.com/tgbv/drawla/ws"
	"goji.io"
)

func main() {
	// load dotenv
	err := dotenv.Load()
	if err != nil {
		panic("Could not load .env file!")
	}

	// first init goji mux
	mux := goji.NewMux()

	// pass the mux to the router as reference
	// the routes will be binded with controllers
	routes.Init(mux)

	// listen HTTP server
	go (func() {
		uri := os.Getenv("L_HOST") + ":" + os.Getenv("L_PORT")

		error := http.ListenAndServe(uri, mux)
		if error != nil {
			panic(error)
		}
	})()
	log.Println("HTTP server started listening on port:", os.Getenv("L_PORT"))

	// listen ws server
	go ws.Init()
	log.Println("WS server started listening on port:", os.Getenv("L_WS_PORT"))

	// waiter
	var input string
	fmt.Println("Press enter to exit program.")
	fmt.Scanln(&input)
}
