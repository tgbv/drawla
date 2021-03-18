package ws

import "github.com/ambelovsky/gosf"

/*
*	handle drawing signal from client
 */
func draw(c *gosf.Client, r *gosf.Request) *gosf.Message {

	m := new(gosf.Message)
	m.Success = true
	m.Body = r.Message.Body

	gosf.Broadcast(DRAWING_ROOM, "draw", m)

	return nil
}
