# Mars Rover as a Go Practice

[![test status](https://github.com/IVIR3zaM/GoPractices-MarsRover/actions/workflows/tests.yml/badge.svg?branch=main "test status")](https://github.com/IVIR3zaM/GoPractices-MarsRover/actions)

On this repository, I tried to put my Golang knowledge into action and solve a sample problem.

## Problem Description
You are part of the team that explores Mars by sending remotely controlled vehicles to the surface of the planet. Write an idiomatic piece of software that translates the commands sent from earth to actions executed by the rover yielding a final state.

When the rover touches down on Mars, it is initialized with its current coordinates and the direction it is facing. These could be any coordinates, supplied as a string in the format `(x, y, direction)`, e.g. `(4, 2, EAST)`.

Once initialized, the rover should be able to report its coordinates in the format  `(4, 2) EAST`


The rover is given a command string that contains multiple commands. This string must then be broken into each command and then executed. Implement the following commands: 

`F -> Move forward on the current heading`

`B -> Move backward on the current heading`

`L -> Rotate left by 90 degrees`

`R -> Rotate right by 90 degrees`

An example command might be `FLFFFRFLB`

Once the full command string has been followed, the rover reports its current coordinates and heading in the format `(6, 4) NORTH`


Previous missions have had to be aborted due to obstacles that caused damage to the rover. Given a set of coordinates for all the known obstacles in the format:

`[[1,4], [3,5], [7,4], [6,5]]`

When the rover would enter a coordinate with an obstacle, instead stop at the coordinate immediately before and report position, heading and Stopped due to collision, e.g. `(5, 5) EAST STOPPED`
