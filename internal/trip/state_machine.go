package trip

import (
	"errors"
	"time"
)

type TripState string

const (
	Created TripState = "CREATED"
	Started TripState = "STARTED"
	Active  TripState = "ACTIVE"
	Ended   TripState = "ENDED"
)

type Trip struct {
	ID        string
	State     TripState
	UpdatedAt time.Time
}

type StateMachine struct {
	trip *Trip
}

func NewStateMachine(trip *Trip) *StateMachine {
	return &StateMachine{trip: trip}
}

func (sm *StateMachine) StartTrip() error {
	if sm.trip.State != Created {
		return errors.New("trip cannot be started from current state")
	}
	sm.trip.State = Started
	sm.trip.UpdatedAt = time.Now()
	return nil
}

func (sm *StateMachine) ActivateTrip() error {
	if sm.trip.State != Started {
		return errors.New("trip cannot be activated from current state")
	}
	sm.trip.State = Active
	sm.trip.UpdatedAt = time.Now()
	return nil
}

func (sm *StateMachine) EndTrip() error {
	if sm.trip.State != Active {
		return errors.New("trip cannot be ended from current state")
	}
	sm.trip.State = Ended
	sm.trip.UpdatedAt = time.Now()
	return nil
}

func (sm *StateMachine) GetCurrentState() TripState {
	return sm.trip.State
}
