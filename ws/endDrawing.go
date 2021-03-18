package ws

import "github.com/ambelovsky/gosf"

/*
*	handle end drawing signal from client
 */
func endDrawing(c *gosf.Client, r *gosf.Request) *gosf.Message {
	gosf.Broadcast(DRAWING_ROOM, "stopDrawing", r.Message)

	return nil
}

/*
*	alias to func above
 */
func stopDrawing(c *gosf.Client, r *gosf.Request) *gosf.Message {
	return endDrawing(c, r)
}
