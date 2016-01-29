package robot

//directions that things can be facing
const (
	N = iota
	E = iota
	S = iota
	W = iota
)

/*Action3 are a list of function that a robot can preform.*/
type Action3 struct {
	Name       string
	ActionList []Action
}

/*Robot3 creats a robot for step 3 of the unit tests.*/
func Robot3(name, script string, actions chan Action3, log chan string) {
	var actList []Action
	cmd := make(chan Command)
	acts := make(chan Action)
	go Robot(cmd, acts)
	for _, char := range script {
		cmd <- Command(char)
		if act := <-acts; act != nil {
			actList = append(actList, act)
		} else {
			log <- "Robot " + name + " recieved an invalid command: " + string(char)
			break
		}
	}
	close(cmd)
	actions <- Action3{name, actList}
}

//state state of a robot
type state struct {
	pos    *DirAt
	active bool
}

/*Room3 runs the colision and orchestration of multiple robots.*/
func Room3(extent Rect, robots []Place, actions chan Action3, report chan []Place, log chan string) {
	//Setup robot tracking
	robotStates := make(map[string]*state)
	for r, robot := range robots {
		_, ok := robotStates[robot.Name]
		switch {
		case robot.Name == "":
			log <- "A robot does not have a name"
		case ok:
			log <- "Robot " + robot.Name + " is listed twice"
		case robotOnSquare(robot.DirAt.Pos, robotStates) != "":
			log <- "Robot " + robot.Name + " was placed on an occupied space"
		case !inRoom(robot.DirAt.Pos, extent):
			log <- "Robot " + robot.Name + " was placed outside the room"
		default:
			robotStates[robot.Name] = &state{&(robots[r].DirAt), true}
			continue
		}
		report <- robots
		return
	}

	//Run the commands
	clearState := getState()
	for activeRobots(robotStates) {
		robot := <-actions
		robotState, ok := robotStates[robot.Name]
		if !ok {
			log <- "Robot " + robot.Name + " doesn't exist"
			break
		}
		setState(*robotState.pos)
		for _, action := range robot.ActionList {
			action()
			otherRobot := robotOnSquare(Pos{RU(X), RU(Y)}, robotStates)
			switch {
			case !inRoom(Pos{RU(X), RU(Y)}, extent):
				log <- "Robot " + robot.Name + " bumped into the wall"
			case otherRobot != robot.Name && otherRobot != "":
				log <- "Robot " + robot.Name + " bumped into " + otherRobot
			default:
				*robotState.pos = getState()
			}
			setState(*robotState.pos)
		}
		robotStates[robot.Name].active = false
	}
	report <- robots
	setState(clearState)
}

/*activeRobots checks if there is still an active robot.*/
func activeRobots(robots map[string]*state) bool {
	for _, state := range robots {
		if state.active {
			return true
		}
	}
	return false
}

/*robotOnSquare checks if there is a robot in a give position.*/
func robotOnSquare(square Pos, otherRobots map[string]*state) string {
	for name, state := range otherRobots {
		if square == state.pos.Pos {
			return name
		}
	}
	return ""
}

/*====================STEP 2====================*/

/*Action are function that a robot can preform.*/
type Action func()

/*Robot creats a robot for step 2.*/
func Robot(commands chan Command, actions chan Action) {
	for command, ok := <-commands; ok; command, ok = <-commands {
		switch command {
		case 'A':
			actions <- Advance
		case 'L':
			actions <- Left
		case 'R':
			actions <- Right
		default:
			actions <- nil
		}
	}
	close(actions)
}

/*Room creates a room for a robot to operate in.*/
func Room(extent Rect, place DirAt, actions chan Action, places chan DirAt) {
	cleanState := getState()
	setState(place)
	for action, ok := <-actions; ok; action, ok = <-actions {
		oldState := getState()
		action()
		if !inRoom(Pos{RU(X), RU(Y)}, extent) {
			setState(oldState)
		}
	}
	places <- getState()
	setState(cleanState)
}

/*inRoom checks if a position is in a room.*/
func inRoom(point Pos, room Rect) bool {
	minX, minY := room.Min.Easting, room.Min.Northing
	maxX, maxY := room.Max.Easting, room.Max.Northing
	x, y := point.Easting, point.Northing
	return minX <= x && x <= maxX && minY <= y && y <= maxY
}

/*setState sets the global variables.*/
func setState(state DirAt) {
	Facing, X, Y = state.Dir, int(state.Pos.Easting), int(state.Pos.Northing)
}

/*getState gets the global variables.*/
func getState() DirAt {
	return DirAt{Facing, Pos{RU(X), RU(Y)}}
}

/*====================STEP 1====================*/
/*String representaion heading as a string.*/
func (d Dir) String() string {
	switch d {
	case N:
		return "north"
	case E:
		return "east"
	case S:
		return "south"
	case W:
		return "west"
	default:
		return ""
	}
}

/*Advance moves the robot forward.*/
func Advance() {
	switch Facing {
	case N:
		Y++
	case S:
		Y--
	case E:
		X++
	case W:
		X--
	}
}

/*Right turns the robot to the right.*/
func Right() {
	Facing = (Facing + 1) % 4
}

/*Left turns the robot to the left.*/
func Left() {
	Facing = (Facing + 3) % 4
}
