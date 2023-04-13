package main

import (
	"FlightTracker/cmd/server"
	"FlightTracker/domain/config"
	"context"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

const (
	appName     = "flighttracker"
	description = "FlightTracker computes flight information for multiple user routes"
)

func main() {
	conf := config.NewFlightTrackerServiceConfig()

	// set the appName, description and config in the admin

	// set logger

	ctx := context.Background()

	// Start service via AdminService
	service, err := server.NewFlightTrackerService(ctx, conf)
	if err != nil {
		log.Fatal("unable to create service")
	}

	// start admin service using ctx, appName, service

	flightCalculatorAPI := service.NewFlightCalculatorAPI()
	restful.DefaultContainer.Add(flightCalculatorAPI)

	// add swagger and api doc
	//config := restfulspec.Config{
	//	WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
	//	APIPath:                       "/apidocs.json",
	//	PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	//restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	//
	//// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	//// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	//// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	//http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("Projects/swagger-ui/dist"))))

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	//
	// admin service start using ctx and the webservices added
}
