package model

const (
	Layer1 = 0
	Layer2 = 1
)

type Element struct {
	Layer int
	Char  string
}

func (e *Element) Rune() rune {
	return []rune(e.Char)[0]
}

func GetElement(r rune) Element {
	for _, e := range Elements {
		if e.Rune() == r {
			return e
		}
	}
	return Empty

}

var (
	Empty = Element{Layer2, "-"}
	Floor = Element{Layer1, "."}

	// walls
	AngleInLeft        = Element{Layer1, "╔"}
	WallFront          = Element{Layer1, "═"}
	AngleInRight       = Element{Layer1, "┐"}
	WallRight          = Element{Layer1, "│"}
	AngleBackRight     = Element{Layer1, "┘"}
	WallBack           = Element{Layer1, "─"}
	AngleBackLeft      = Element{Layer1, "└"}
	WallLeft           = Element{Layer1, "║"}
	WallBackAngleLeft  = Element{Layer1, "┌"}
	WallBackAngleRight = Element{Layer1, "╗"}
	AngleOutRight      = Element{Layer1, "╝"}
	AngleOutLeft       = Element{Layer1, "╚"}
	Space              = Element{Layer1, " "}

	// laser machine
	LaserMachineChargingLeft  = Element{Layer1, "˂"}
	LaserMachineChargingRight = Element{Layer1, "˃"}
	LaserMachineChargingUp    = Element{Layer1, "˄"}
	LaserMachineChargingDown  = Element{Layer1, "˅"}

	// lase machine ready
	LaserMachineReadyLeft  = Element{Layer1, "◄"}
	LaserMachineReadyRight = Element{Layer1, "►"}
	LaserMachineReadyUp    = Element{Layer1, "▲"}
	LaserMachineReadyDown  = Element{Layer1, "▼"}

	// other stuff
	Start       = Element{Layer1, "S"}
	Exit        = Element{Layer1, "E"}
	Hole        = Element{Layer1, "O"}
	Box         = Element{Layer2, "B"}
	ZombieStart = Element{Layer1, "Z"}
	Gold        = Element{Layer1, "$"}

	// your robot
	Robo            = Element{Layer2, "☺"}
	RoboFalling     = Element{Layer2, "o"}
	RoboFlying      = Element{Layer2, "*"}
	RoboFlyingOnBox = Element{Layer2, "№"}
	RoboLaser       = Element{Layer2, "☻"}

	// other robot
	RoboOther            = Element{Layer2, "X"}
	RoboOtherFalling     = Element{Layer2, "x"}
	RoboOtherFlying      = Element{Layer2, "^"}
	RoboOtherFlyingOnBox = Element{Layer2, "%"}
	RoboOtherLaser       = Element{Layer2, "&"}

	// laser
	LaserLeft  = Element{Layer2, "←"}
	LaserRight = Element{Layer2, "→"}
	LaserUp    = Element{Layer2, "↑"}
	LaserDown  = Element{Layer2, "↓"}

	// zombie
	FemaleZombie = Element{Layer2, "♀"}
	MaleZombie   = Element{Layer2, "♂"}
	ZombieDie    = Element{Layer2, "✝"}

	// system elements, don"t touch it
	Fog        = Element{Layer1, "F"}
	Background = Element{Layer2, "G"}

	// All elements in slice
	Elements = []Element{
		Empty,
		Floor,

		// walls
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
		Space,

		// laser machine
		LaserMachineChargingLeft,
		LaserMachineChargingRight,
		LaserMachineChargingUp,
		LaserMachineChargingDown,

		// lase machine ready
		LaserMachineReadyLeft,
		LaserMachineReadyRight,
		LaserMachineReadyUp,
		LaserMachineReadyDown,

		// other stuff
		Start,
		Exit,
		Hole,
		Box,
		ZombieStart,
		Gold,

		// your robot
		Robo,
		RoboFalling,
		RoboFlying,
		RoboFlyingOnBox,
		RoboLaser,

		// other robot
		RoboOther,
		RoboOtherFalling,
		RoboOtherFlying,
		RoboOtherFlyingOnBox,
		RoboOtherLaser,

		// laser
		LaserLeft,
		LaserRight,
		LaserUp,
		LaserDown,

		// zombie
		FemaleZombie,
		MaleZombie,
		ZombieDie,

		// system elements, don"t touch it
		Fog,
		Background,
	}
)
