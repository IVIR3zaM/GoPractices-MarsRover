package rover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoverTestSuite struct {
	suite.Suite
	rover *Rover
}

func (s *RoverTestSuite) SetupTest() {
	if rover, err := NewRover("(4, 2, EAST)", nil); err == nil {
		s.rover = rover
	} else {
		s.Fail(fmt.Sprintf("Can not initialize rover, err: %s", err))
	}
}

func (s *RoverTestSuite) TestRoverOutput() {
	s.Equal("(4, 2) EAST", s.rover.Output())
}

func (s *RoverTestSuite) TestMoveForward() {
	s.executeCmd("F", "(5, 2) EAST")

	s.rover.SetDirection(South)
	s.executeCmd("F", "(5, 1) SOUTH")

	s.rover.SetDirection(West)
	s.executeCmd("F", "(4, 1) WEST")

	s.rover.SetDirection(North)
	s.executeCmd("F", "(4, 2) NORTH")
}

func (s *RoverTestSuite) TestMoveBackward() {
	s.executeCmd("B", "(3, 2) EAST")

	s.rover.SetDirection(South)
	s.executeCmd("B", "(3, 3) SOUTH")

	s.rover.SetDirection(West)
	s.executeCmd("B", "(4, 3) WEST")

	s.rover.SetDirection(North)
	s.executeCmd("B", "(4, 2) NORTH")
}

func (s *RoverTestSuite) TestTurnRight() {
	s.executeCmd("R", "(4, 2) SOUTH")
	s.executeCmd("R", "(4, 2) WEST")
	s.executeCmd("R", "(4, 2) NORTH")
	s.executeCmd("R", "(4, 2) EAST")
}

func (s *RoverTestSuite) TestTurnLeft() {
	s.executeCmd("L", "(4, 2) NORTH")
	s.executeCmd("L", "(4, 2) WEST")
	s.executeCmd("L", "(4, 2) SOUTH")
	s.executeCmd("L", "(4, 2) EAST")
}

func (s *RoverTestSuite) TestFullCommand() {
	s.executeCmd("FLFFFRFLB", "(6, 4) NORTH")
}

func (s *RoverTestSuite) TestFullCommandWithObstacles() {
	s.rover.SetObstacles("[[1,4], [3,5], [7,4], [6,5]]")
	s.executeCmd("FLFFFRFLB", "(5, 5) EAST STOPPED")
}

func (s *RoverTestSuite) executeCmd(cmd string, output string) {
	err := s.rover.Execute(cmd)
	s.Nil(err)
	s.Equal(output, s.rover.Output())
}

func TestRover(t *testing.T) {
	suite.Run(t, new(RoverTestSuite))
}
