package model

const (
	Nothing = ""
	Reset   = "ACT(0)"
	Jump    = "ACT(1)"
	JumpTo  = "ACT(1),"
	PullTo  = "ACT(2),"

	Up    Direction = "UP"
	Right Direction = "RIGHT"
	Down  Direction = "DOWN"
	Left  Direction = "LEFT"
)

var (
	Directions     = []Direction{Up, Right, Down, Left}
	DirectionsMaps = map[Direction]Point{
		Up:    {0, -1},
		Right: {1, 0},
		Down:  {0, 1},
		Left:  {-1, 0},
	}
)

type Direction string

func (d Direction) index() int {
	for i, dir := range Directions {
		if dir == d {
			return i
		}
	}
	return -1
}

func (d Direction) Clockwise() Direction {
	return Directions[(d.index()+1)%len(Directions)]
}

func (d Direction) Invert() Direction {
	return d.Clockwise().Clockwise()
}

func (d Direction) Change(p Point) Point {
	next := DirectionsMaps[d]
	return Point{p.X + next.X, p.Y + next.Y}
}

func (d Direction) ChangeX(x int) int {
	p := Point{x, 0}
	return d.Change(p).X
}
func (d Direction) ChangeY(y int) int {
	p := Point{0, y}
	return d.Change(p).Y
}
