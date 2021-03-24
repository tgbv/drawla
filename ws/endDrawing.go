package ws

import (
	"github.com/ambelovsky/gosf"
	"github.com/tgbv/drawla/state"
)

/*
*	handle end drawing signal from client
 */
func endDrawing(c *gosf.Client, r *gosf.Request) *gosf.Message {
	rid := r.Message.Text

	// if room does not exist, disconnect without a word
	if state.Rooms[rid] == nil {
		c.Disconnect()
	}

	gosf.Broadcast(rid, "stopDrawing", r.Message)

	return nil
}

/*
*	alias to func above
 */
func stopDrawing(c *gosf.Client, r *gosf.Request) *gosf.Message {
	return endDrawing(c, r)
}
