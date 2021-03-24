package lib

/*
	Contains state-related data structures and methods
*/

/*
*	a point from rooms structure
 */
type Point struct {
	Coords struct {
		X int
		Y int
	}
}

/*
*	holds a room
	syntax: {
		points: [
			{
				coords: {
					x: (int)
					y: (int)
				}
			}
		],

		...
	}
*/
type Room struct {
	Id     string
	Points []Point
}

/*
*	Holds THE rooms
	syntax: room_id: Room
*/
type Rooms map[string]*Room

/*
*	adds a new Room to Rooms map
 */
func (r *Rooms) Add(id string) (*Rooms, *Room) {
	(*r)[id] = &Room{
		Id:     id,
		Points: make([]Point, 0),
	}

	return r, (*r)[id]
}

/*
*	adds a new point to a Room
 */
func (r *Room) AddPoint(x int, y int) (*Room, *Point) {
	p := Point{
		Coords: struct {
			X int
			Y int
		}{x, y},
	}

	r.Points = append(r.Points, p)

	return r, &p
}

/*
*	deletes a point from Room based on callback
	O(n)
*/
func (r *Room) DeletePoint(f func(Point) bool) (*Room, bool) {

	n := len(r.Points)

	for i, v := range r.Points {
		if f(v) == true {
			r.Points[i] = r.Points[n-1]
			r.Points = r.Points[:n-1]

			return r, true
		}
	}

	return r, false
}
