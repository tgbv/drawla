package ws

import (
	"github.com/ambelovsky/gosf"
	"github.com/tgbv/drawla/state"
)

/*
*	handle start drawing signal from client
 */
func startDrawing(c *gosf.Client, r *gosf.Request) *gosf.Message {
	rid := r.Message.Text

	// if room does not exist, disconnect without a word
	if state.Rooms[rid] == nil {
		c.Disconnect()
	}

	gosf.Broadcast(rid, "startDrawing", r.Message)

	return nil
}
