package handlers

import (
	"FlightTracker/domain/entities"
	domain "FlightTracker/domain/usecases"
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

const (
	apiVersion = "v1"
)

type flightCalculatorAPI struct {
	calculatorManager domain.FlightCalculatorManager
}

// SetupFlightCalculatorAPI register the RESTful API to interact with the flight calculator for passenger routes
func SetupFlightCalculatorAPI(ctx context.Context, calculatorManager domain.FlightCalculatorManager) *restful.WebService {
	fc := &flightCalculatorAPI{
		calculatorManager: calculatorManager,
	}
	ws := &restful.WebService{}

	ws.Route(ws.POST("/calculate/v1").
		To(fc.handleCalculate)).
		ApiVersion(apiVersion).
		Doc("Takes a list of departure and arrival airport codes representing the flights of a persons trip and returns " +
			"the initial location of the person and the final destination").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	return ws
}

func (fc *flightCalculatorAPI) handleCalculate(request *restful.Request, response *restful.Response) {
	// get context from request

	result := entities.Flight{}
	defer func() {
		if err := response.WriteAsJson(CalculateResponseV1{
			Resources: result,
		}); err != nil {
			// log error
		}
	}()

	inputFlights := &CalculateRequest{}
	if err := request.ReadEntity(&inputFlights.Flights); err != nil {
		// add MSA errors for the response
		// add logging
		fmt.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println(inputFlights)
	}

	if resultFlight, err := fc.calculatorManager.Calculate(inputFlights.Flights); err != nil {
		// add MSA errors for the response
		// add logging
		response.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		result = resultFlight
	}
}
