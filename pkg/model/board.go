package model

const (
	Size = 20
)

type Progress struct {
	Total      int `json:"total"`
	Current    int `json:"current"`
	LastPassed int `json:"lastPassed"`
}

type Board struct {
	ShowName      bool     `json:"showName"`
	Offset        Point    `json:"offset"`
	Layers        []string `json:"layers"`
	LevelProgress Progress `json:"level_progress"`
}

// Movement functions
func (b Board) DoNothing() string {
	return Nothing
}

func (b Board) Reset() string {
	return Reset
}

func (b Board) Go(d Direction) string {
	return string(d)
}

func (b Board) Jump() string {
	return Jump
}

func (b Board) JumpTo(d Direction) string {
	return JumpTo + string(d)
}

func (b Board) PullTo(d Direction) string {
	return PullTo + string(d)
}

// Board functions
func (b Board) get(layer int, elems ...Element) []Point {
	pts := make([]Point, 0)
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			p := Point{x, y}
			if b.IsAt(layer, p, elems...) {
				pts = append(pts, p)
			}
		}
	}
	return pts
}

func (b Board) IsAt(layer int, p Point, elems ...Element) bool {
	for _, e := range elems {
		if e.Rune() == []rune(b.Layers[layer])[p.Y*Size+p.X] {
			return true
		}
	}
	return false
}

func (b Board) GetMe() Point {
	pts := b.get(Layer2,
		RoboFalling,
		RoboFlying,
		RoboFlyingOnBox,
		RoboLaser,
		Robo)
	return pts[0]
}

// Returns list of coordinates for all visible enemy Robots.
func (b Board) GetOtherHeroes() []Point {
	return b.get(Layer2,
		RoboOtherFalling,
		RoboOtherFlying,
		RoboOtherFlyingOnBox,
		RoboOtherLaser,
		RoboOther)
}

// Returns list of coordinates for all visible LaserMachines.
func (b Board) GetLaserMachines() []Point {
	return b.get(Layer1,
		LaserMachineChargingLeft,
		LaserMachineChargingRight,
		LaserMachineChargingUp,
		LaserMachineChargingDown,

		LaserMachineReadyLeft,
		LaserMachineReadyRight,
		LaserMachineReadyUp,
		LaserMachineReadyDown)
}

// Returns list of coordinates for all visible Lasers.
func (b Board) GetLasers() []Point {
	return b.get(Layer2,
		LaserLeft,
		LaserRight,
		LaserUp,
		LaserDown)
}

// Returns list of coordinates for all visible Walls.
func (b Board) getWalls() []Point {
	return b.get(Layer1,
		AngleInLeft,
		WallFront,
		AngleInRight,
		WallRight,
		AngleBackRight,
		WallBack,
		AngleBackLeft,
		WallLeft,
		WallBackAngleLeft,
		WallBackAngleRight,
		AngleOutRight,
		AngleOutLeft,
		Space)
}

// Returns list of coordinates for all visible Boxes.
func (b Board) GetBoxes() []Point {
	return b.get(Layer2,
		Box,
		RoboFlyingOnBox,
		RoboOtherFlyingOnBox)
}

// Returns list of coordinates for all visible Holes.
func (b Board) GetHoles() []Point {
	return b.get(Layer1,
		Hole,
		RoboFalling,
		RoboOtherFalling)
}

// Returns list of coordinates for all visible Exit points.
func (b Board) GetExits() []Point {
	return b.get(Layer1, Exit)
}

// Returns list of coordinates for all visible Gold.
func (b Board) GetGold() []Point {
	return b.get(Layer1, Gold)
}

// Checks if your robot is alive.
func (b Board) IsMeAlive() bool {
	collisions := b.get(Layer2, RoboFalling, RoboLaser)
	return len(collisions) == 0
}

// Is it possible to go through the cell with {x,y} coordinates.
func (b Board) IsBarrierAt(p Point) bool {
	return !b.IsAt(Layer1, p, Floor, Start, Exit, Gold, Hole) || !b.IsAt(Layer2, p, Empty, Gold,
		LaserDown, LaserUp, LaserLeft, LaserRight,
		RoboOther, RoboOtherFlying, RoboOtherFalling, RoboOtherLaser,
		Robo, RoboFlying, RoboFalling, RoboLaser)
}

// Get element at point
func (b Board) GetAt(layer int, p Point) Element {
	return GetElement([]rune(b.Layers[layer])[p.Y*Size+p.X])
}

// Returns element at position specified
func (b Board) GetNear(layer int, p Point) []Element {
	elems := make([]Element, 0)
	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			if !b.OutOfSize(Point{x, y}) && (x != p.X || y != p.Y) {
				elems = append(elems, b.GetAt(layer, p))
			}
		}
	}
	return elems
}

// Count near elements
func (b Board) CountNear(layer int, p Point, element Element) int {
	cnt := 0
	for _, e := range b.GetNear(layer, p) {
		if e.Rune() == element.Rune() {
			cnt++
		}
	}
	return cnt
}

// Check if is near element
func (b Board) IsNear(layer int, p Point, element Element) bool {
	if !b.OutOfSize(p) {
		return b.CountNear(layer, p, element) > 0
	}
	return false
}

func (b Board) OutOfSize(p Point) bool {
	return !(p.X > 0 && p.X < Size && p.Y > 0 && p.Y < Size)
}
