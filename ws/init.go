package ws

import (
	"os"
	"strconv"

	"github.com/ambelovsky/gosf"
)

const DRAWING_ROOM string = ""

/*
*	initialize ws server
 */
func Init() {

	// start drawing
	gosf.Listen("startDrawing", startDrawing)

	// draw
	gosf.Listen("draw", draw)

	// stop drawing
	gosf.Listen("stopDrawing", endDrawing)

	// get ws port
	p, err := strconv.Atoi(os.Getenv("L_WS_PORT"))
	if err != nil {
		panic("ENV error occurred!")
	}

	// start ws server
	gosf.Startup(map[string]interface{}{
		"port": p,
		"host": os.Getenv("L_WS_HOST"),
	})

}
