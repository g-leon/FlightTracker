package usecases

import "FlightTracker/domain/entities"

// FlightCalculatorManager is a manager for calculating flights of a passenger
type FlightCalculatorManager interface {
	Calculate([]entities.Flight) (entities.Flight, error)
}
