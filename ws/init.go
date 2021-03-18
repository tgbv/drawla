package ws

import "github.com/ambelovsky/gosf"

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

	// start ws server
	gosf.Startup(map[string]interface{}{"port": 9999})

}
