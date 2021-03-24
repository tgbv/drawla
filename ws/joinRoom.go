package ws

import (
	"github.com/ambelovsky/gosf"
	"github.com/tgbv/drawla/state"
)

/*
*	joins client to custom room
 */
func joinRoom(c *gosf.Client, r *gosf.Request) *gosf.Message {
	rid := r.Message.Text

	// if room does not exist, disconnect without a word
	if state.Rooms[rid] == nil {
		c.Disconnect()
		return nil
	}

	c.Join(rid)

	m := new(gosf.Message)
	m.Success = true
	m.Text = rid

	return m
}
