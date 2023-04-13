package handlers

import "FlightTracker/domain/entities"

// CalculateRequest encapsulates the input used for the flight calculator API
// Used for marshaling input
type CalculateRequest struct {
	Flights []entities.Flight `json:""`
}

// CalculateResponseV1 encapsulates the flight calculate response
type CalculateResponseV1 struct {
	// MsaMetaInfo
	// MsaErrors
	Resources []string `json:""`
}
