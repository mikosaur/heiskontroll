package elevator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTravelToUnknownFloor(t *testing.T) {
	minFloor := 1
	maxFloor := 10
	timePerFloor := 2
	timeToOpenDoor := 1
	elevator := NewElevator(minFloor, maxFloor, timePerFloor, timeToOpenDoor)

	err := elevator.GoToFloor(11)
	require.NotNil(t, err)

	unknownFloorError := &UnknownFloorError{}
	require.ErrorAs(t, err, &unknownFloorError)
}

func TestTravelTime(t *testing.T) {

	minFloor := 1
	maxFloor := 10
	timePerFloor := 2
	timeToOpenDoor := 1
	elevator := NewElevator(minFloor, maxFloor, timePerFloor, timeToOpenDoor)

	time := elevator.GetTimeToFloorInSeconds(10)

	// From floor 1 to 10 is 9 floors.
	// Total 9 * 2 = 18 seconds to travel all floors
	// Total of 2 seconds to open and close the doors
	require.Equal(t, 20, time)
}

func TestDirection(t *testing.T) {
	minFloor := 1
	maxFloor := 10
	timePerFloor := 2
	timeToOpenDoor := 1
	elevator := NewElevator(minFloor, maxFloor, timePerFloor, timeToOpenDoor)

	dir := elevator.GetDirection()
	require.Equal(t, DirectionUp, dir)

	err := elevator.GoToFloor(9)
	require.NoError(t, err)

	dir = elevator.GetDirection()
	require.Equal(t, DirectionUp, dir)

	err = elevator.GoToFloor(2)
	require.NoError(t, err)

	dir = elevator.GetDirection()
	require.Equal(t, DirectionDown, dir)

	err = elevator.GoToFloor(1)
	require.NoError(t, err)

	dir = elevator.GetDirection()
	require.Equal(t, DirectionUp, dir)

	err = elevator.GoToFloor(10)
	require.NoError(t, err)

	dir = elevator.GetDirection()
	require.Equal(t, DirectionDown, dir)

}

func TestEmergency(t *testing.T) {

	minFloor := 1
	maxFloor := 10
	timePerFloor := 2
	timeToOpenDoor := 1
	elevator := NewElevator(minFloor, maxFloor, timePerFloor, timeToOpenDoor)

	elevator.SetEmergency(true)

	err := elevator.GoToFloor(5)
	require.NotNil(t, err)
	emergencyErr := &IsEmergencyError{}
	require.ErrorAs(t, err, &emergencyErr)
}
