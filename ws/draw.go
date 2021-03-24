package ws

import (
	"github.com/ambelovsky/gosf"
	"github.com/tgbv/drawla/state"
)

/*
*	handle drawing signal from client
 */
func draw(c *gosf.Client, r *gosf.Request) *gosf.Message {
	rid := r.Message.Text

	// if room does not exist, disconnect without a word
	if state.Rooms[rid] == nil {
		c.Disconnect()
	}

	m := new(gosf.Message)
	m.Success = true
	m.Body = r.Message.Body

	gosf.Broadcast(rid, "draw", m)

	return nil
}
