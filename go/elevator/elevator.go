package elevator

import (
	"fmt"
	"math"
)

type Direction string

const (
	DirectionUp   Direction = "UP"
	DirectionDown Direction = "DOWN"
)

type Elevator struct {
	currentFloor int // Where the elevator is currently located

	maxFloor int // Highest floor
	minFloor int // Lowest floor, can be negative

	timePerFloor   int // In seconds
	timeToOpenDoor int // In seconds

	direction Direction // the direction of the elevator

	isEmergency bool // If true, the elevator is stopped and wont move until emergency has been set to false
}

func NewElevator(minFloor, maxFloor, timePerFloor, timeToOpenDoor int) *Elevator {
	return &Elevator{
		currentFloor:   minFloor,
		minFloor:       minFloor,
		maxFloor:       maxFloor,
		timePerFloor:   timePerFloor,
		timeToOpenDoor: timeToOpenDoor,
		direction:      DirectionUp,
		isEmergency:    false,
	}
}

func (e *Elevator) GetCurrentFloor() int {
	return e.currentFloor
}

func (e *Elevator) GetDirection() Direction {
	return e.direction
}

func (e *Elevator) GetTimeToFloorInSeconds(toFloor int) int {
	// *2 is for close door at origin floor, and open at destination floor
	doorTime := e.timeToOpenDoor * 2

	floorsToTravel := math.Abs(float64(e.currentFloor) - float64(toFloor))
	travelTime := doorTime + int(floorsToTravel)*e.timePerFloor

	return travelTime
}

func (e *Elevator) GoToFloor(toFloor int) error {
	// Check if the caller is attempting to go to a floor that we do not support

	if e.isEmergency {
		return &IsEmergencyError{}
	}

	if toFloor < e.minFloor {
		return NewUnkownFloorError(fmt.Sprintf("lowest available floor is: %d", e.minFloor))
	}

	if toFloor > e.maxFloor {
		return NewUnkownFloorError(fmt.Sprintf("highest available floor is %d", e.maxFloor))
	}

	dir := e.currentFloor - toFloor
	var nextDir Direction
	if toFloor == e.maxFloor {
		nextDir = DirectionDown
	}

	if toFloor == e.minFloor {
		nextDir = DirectionUp
	}

	if nextDir == "" {
		if dir < 0 {
			nextDir = DirectionUp
		} else {
			nextDir = DirectionDown
		}
	}

	e.direction = nextDir

	// Move to wanted floor
	// time.Sleep(time.Duration(e.GetTimeToFloorInSeconds(toFloor)))
	e.currentFloor = toFloor

	return nil
}

func (e *Elevator) SetEmergency(emergency bool) {
	e.isEmergency = emergency
}
