package elevator

type UnknownFloorError struct {
	Err string
}

func (e *UnknownFloorError) Error() string {
	return e.Err
}

func NewUnkownFloorError(err string) *UnknownFloorError {
	return &UnknownFloorError{Err: err}
}

type IsEmergencyError struct{}

func (e *IsEmergencyError) Error() string {
	return ""
}
