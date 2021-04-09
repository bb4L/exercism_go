package robot

// step1
const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) String() string {
	var direction string
	switch d {
	case N:
		direction = "North"
	case S:
		direction = "South"
	case E:
		direction = "East"
	case W:
		direction = "West"
	}
	return direction
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case S:
		Step1Robot.Y -= 1
	case E:
		Step1Robot.X += 1
	case W:
		Step1Robot.X -= 1
	}
}

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir - 1 + 4) % 4
}

// step2

type Action byte

func StartRobot(cmd chan Command, act chan Action) {
	go func() {
		defer close(act)

		for c := range cmd {
			act <- Action(c)
		}
	}()
}

func Room(rect Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	go func() {
		defer close(rep)
		for a := range act {
			robot.DoAction(a, rect)
		}
		rep <- robot
	}()
}

func (robot *Step2Robot) DoAction(a Action, rect Rect) {
	switch a {
	case 'L':
		robot.Dir = (robot.Dir + 3) % 4
	case 'R':
		robot.Dir = (robot.Dir + 1) % 4
	case 'A':
		x := robot.Pos.Easting
		y := robot.Pos.Northing
		switch robot.Dir {
		case N:
			y += 1
		case S:
			y -= 1

		case E:
			x += 1
		case W:
			x -= 1

		}
		if (x <= rect.Max.Easting && x >= rect.Min.Easting) && (y <= rect.Max.Northing && y >= rect.Min.Northing) {
			robot.Pos.Northing = y
			robot.Pos.Easting = x
		}
	}
}

// step3

type Action3 struct {
	Action
	Name string
}

func Room3(rect Rect, robots []Step3Robot, act chan Action3, rep chan []Step3Robot, log chan string) {

	hasEnded := 0
	step3Robots := make([]Step3Robot, 0, len(robots))
	positions := make(map[Pos]bool)

	robotMap := make(map[string]Step3Robot)

	for _, robot := range robots {
		if _, ok := robotMap[robot.Name]; ok {
			log <- "duplicate robot"
			rep <- robots
			return
		}
		if _, ok := positions[robot.Step2Robot.Pos]; ok {
			log <- "duplicate init position"
			rep <- robots
			return
		}

		x := robot.Step2Robot.Pos.Easting
		y := robot.Step2Robot.Pos.Northing
		if (x > rect.Max.Easting || x < rect.Min.Easting) || (y > rect.Max.Northing || y < rect.Min.Northing) {
			log <- "init outside of rect"
			rep <- robots
			return
		}

		positions[robot.Step2Robot.Pos] = true
		robotMap[robot.Name] = robot
	}

	for action := range act {
		if action.Action == Action(0) {
			hasEnded += 1
			if hasEnded == len(robots) {
				close(act)
				for name, finalRobot := range robotMap {
					step3Robots = append(step3Robots, Step3Robot{name, Step2Robot{finalRobot.Step2Robot.Dir, finalRobot.Step2Robot.Pos}})
				}
				rep <- step3Robots
				return
			} else {
				continue
			}
		}

		robot, ok := robotMap[action.Name]
		if !ok {
			log <- "cmd for unknown robot"
			rep <- step3Robots
			return
		}

		oldPos := robot.Step2Robot.Pos

		actualAction := action.Action
		if actualAction != 'R' && actualAction != 'L' && actualAction != 'A' {
			log <- "invalid command"
			rep <- robots
			return
		}

		robot.Step2Robot.DoAction(action.Action, rect)
		if oldPos.Northing == robot.Step2Robot.Pos.Northing && oldPos.Easting == robot.Step2Robot.Pos.Easting && action.Action == 'A' {
			robot.Step2Robot.Pos = oldPos
			robotMap[action.Name] = robot
			log <- "bounced from wall"
			continue
		}

		if oldPos != robot.Step2Robot.Pos {
			if _, ok := positions[robot.Step2Robot.Pos]; ok {
				log <- "bumping into other robot"
				robot.Step2Robot.Pos = oldPos
				robotMap[action.Name] = robot
				continue
			}

			delete(positions, oldPos)
			positions[robot.Step2Robot.Pos] = true

		}

		robotMap[action.Name] = robot

	}

}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "Invalid name"
	}

	for _, c := range script {
		action <- Action3{Action(c), name}
	}
	action <- Action3{Action(0), name}
}
