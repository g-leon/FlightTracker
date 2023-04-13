package server

import (
	"FlightTracker/api/handlers"
	"FlightTracker/domain/config"
	domain_usecases "FlightTracker/domain/usecases"
	"FlightTracker/usecases"
	"context"
	"github.com/emicklei/go-restful"
)

// FlightTrackerService represents the flight tracker service
type FlightTrackerService struct {
	config                  *config.FlightTrackerServiceConfig
	ctx                     context.Context
	flightCalculatorManager domain_usecases.FlightCalculatorManager
}

// NewFlightTrackerService returns a new instance of FlightTrackerService
func NewFlightTrackerService(ctx context.Context, config *config.FlightTrackerServiceConfig) (*FlightTrackerService, error) {
	return &FlightTrackerService{
		config:                  config,
		ctx:                     ctx,
		flightCalculatorManager: usecases.NewFlightCalculatorImpl(ctx),
	}, nil
}

// NewFlightCalculatorAPI returns a new instance of FlightCalculatorAPI
func (ft *FlightTrackerService) NewFlightCalculatorAPI() *restful.WebService {
	return handlers.SetupFlightCalculatorAPI(ft.ctx, ft.flightCalculatorManager)
}

// handle graceful shutdown
