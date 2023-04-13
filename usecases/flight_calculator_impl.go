package usecases

import (
	"FlightTracker/domain/entities"
	"context"
)

// FlightCalculatorImpl is the implementation of the flight calculator manager
type FlightCalculatorImpl struct {
}

// NewFlightCalculatorImpl returns a new object of type FlightCalculatorImpl
func NewFlightCalculatorImpl(ctx context.Context) *FlightCalculatorImpl {
	return &FlightCalculatorImpl{}
}

// Calculate accepts a list of flights as an input and returns the departure and arrival of the entire trip
func (fc *FlightCalculatorImpl) Calculate(inputFlights []entities.Flight) (entities.Flight, error) {
	sources := make(map[string]bool)
	destinations := make(map[string]bool)

	for _, flight := range inputFlights {
		sources[flight[0]] = true
		destinations[flight[1]] = true
	}

	startAirport := ""
	endAirport := ""

	for source := range sources {
		if !destinations[source] {
			startAirport = source
			break
		}
	}

	for destination := range destinations {
		if !sources[destination] {
			endAirport = destination
			break
		}
	}

	finalAirports := entities.Flight{startAirport, endAirport}

	return finalAirports, nil
}
