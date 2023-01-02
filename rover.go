package marchrover

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	East  = "EAST"
	South = "SOUTH"
	West  = "WEST"
	North = "NORTH"
)

const (
	Forward  = 'F'
	Backward = 'B'
	Left     = 'L'
	Right    = 'R'
)

type Direction string

func toDirection(str string) (*Direction, error) {
	for _, d := range [...]string{
		East,
		South,
		West,
		North,
	} {
		if strings.EqualFold(d, str) {
			dir := Direction(d)
			return &dir, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("direction of %q not found", str))
}

type Rover struct {
	roverCoordination roverCoordination
	stopped           bool
	obstacles         [][2]int
	movingDelats      map[Direction][2]int
	directionDeltas   map[Direction][2]Direction
}

type roverCoordination struct {
	x, y      int
	direction Direction
}

func NewRover(input string, obstacles [][2]int) (*Rover, error) {
	str := input[1 : len(input)-1]
	l := strings.Split(str, ", ")

	if len(l) != 3 {
		return nil, errors.New(fmt.Sprintf("Invalid initialization command of %q", input))
	}

	x, err := strconv.Atoi(l[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid initialization x: %q, error: %s", l[0], err))
	}

	y, err := strconv.Atoi(l[1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid initialization y: %q, error: %s", l[1], err))
	}
	direction, err := toDirection(l[2])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid initialization y: %q, error: %s", l[2], err))
	}

	if obstacles == nil {
		obstacles = [][2]int{}
	}

	movingDelats := map[Direction][2]int{
		East:  [...]int{1, 0},
		South: [...]int{0, -1},
		West:  [...]int{-1, 0},
		North: [...]int{0, 1},
	}

	directionDeltas := map[Direction][2]Direction{
		East:  [...]Direction{South, North},
		South: [...]Direction{West, East},
		West:  [...]Direction{North, South},
		North: [...]Direction{East, West},
	}

	roverCoordination := roverCoordination{x, y, *direction}

	return &Rover{roverCoordination, false, obstacles, movingDelats, directionDeltas}, nil
}

func (r *Rover) SetObstacles(input string) error {
	return json.Unmarshal([]byte(input), &r.obstacles)
}

func (r *Rover) SetDirection(d Direction) {
	r.roverCoordination.direction = d
}

func (r *Rover) Output() string {
	var extra string
	if r.stopped {
		extra = " STOPPED"
	}
	return fmt.Sprintf("(%d, %d) %s%s", r.roverCoordination.x, r.roverCoordination.y, r.roverCoordination.direction, extra)
}

func (r *Rover) Execute(commands string) error {
	newCoordination := r.roverCoordination
	for i, c := range commands {
		if err := newCoordination.coordinate(c, i, r.movingDelats, r.directionDeltas); err != nil {
			return err
		}

		if r.stopped = newCoordination.isOnObstacle(r.obstacles); r.stopped {
			break
		} else {
			r.roverCoordination = newCoordination
		}
	}

	return nil
}

func (rc *roverCoordination) isOnObstacle(obstacles [][2]int) bool {
	for _, obs := range obstacles {
		if obs[0] == rc.x && obs[1] == rc.y {
			return true
		}
	}
	return false
}

func (rc *roverCoordination) coordinate(command rune, pos int, movingDelats map[Direction][2]int, directionDeltas map[Direction][2]Direction) error {
	switch command {
	case Forward:
		delta, err := rc.getMovingDelta(movingDelats)
		if err != nil {
			return err
		}
		rc.x += delta[0]
		rc.y += delta[1]
	case Backward:
		delta, err := rc.getMovingDelta(movingDelats)
		if err != nil {
			return err
		}
		rc.x -= delta[0]
		rc.y -= delta[1]
	case Right:
		delta, err := rc.getDirectionDelta(directionDeltas)
		if err != nil {
			return err
		}
		rc.direction = delta[0]
	case Left:
		delta, err := rc.getDirectionDelta(directionDeltas)
		if err != nil {
			return err
		}
		rc.direction = delta[1]
	default:
		return errors.New(fmt.Sprintf("Invalid command: %q in position %d", command, pos))
	}

	return nil
}

func (rc *roverCoordination) getMovingDelta(movingDelats map[Direction][2]int) (*[2]int, error) {

	delta, ok := movingDelats[rc.direction]
	if ok {
		return &delta, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid direction of %q is set", rc.direction))
}

func (rc *roverCoordination) getDirectionDelta(directionDeltas map[Direction][2]Direction) (*[2]Direction, error) {
	delta, ok := directionDeltas[rc.direction]
	if ok {
		return &delta, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid direction of %q is set", rc.direction))
}
