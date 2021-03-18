package ws

import "github.com/ambelovsky/gosf"

/*
*	handle start drawing signal from client
 */
func startDrawing(c *gosf.Client, r *gosf.Request) *gosf.Message {
	gosf.Broadcast(DRAWING_ROOM, "startDrawing", r.Message)

	return nil
}
