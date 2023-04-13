# README.md

## 1. FlightTracker
A powerful RESTful web service API to track and manage flight paths.

## 2. Description
FlightTracker is a project that aims to provide a simple yet efficient way to track and manage flight paths. The project makes use of the Go programming language for its performance, concurrency support, and simplicity, allowing us to build a reliable and high-performance API.

The application provides a handler `/calculate/v1` that takes a JSON input representing all the flights the passenger has taken, not in sorted order, like `[["SFO", "EWR"]]`, and returns the first departure and last arrival airport, example: `["SFO", "EWR"]`. The service API helps users understand and track flight paths and can be queried for specific information.

During the development of this project, we faced challenges in handling various edge cases and optimizing the performance. In the future, we plan to extend the functionality to support additional query parameters, live flight tracking, and improved error handling.

## 3. How to Install and Run the Project
To install and run the FlightTracker project on your local machine, follow these steps:

1. Ensure that you have Go installed (version 1.16 or higher). If not, follow the installation instructions at `https://golang.org/doc/install`
2. Clone the repository: `git clone https://github.com/g-leon/FlightTracker.git`
3. Navigate to the project directory: `cd FlightTracker`
4. Install the required dependencies `go mod download`
5. Build the project: `go build`
6. Run the FlightTracker API: `./FlightTracker`

The service will now be listening on port 8080.

## 4. How to Use the Project
To use the FlightTracker API, you can send a POST request to the `/calculate/v1` endpoint with the appropriate JSON input. Here's an example using `curl`:
`curl -X POST -H "Content-Type: application/json" -d '[["SFO", "EWR"]]' http://localhost:8080/calculate/v1`

This will return a JSON response with the calculated flight path:
```json
["SFO", "EWR"]
```

## 5. License
This project is licensed under the GNU General Public License v3.0, which allows you to modify and use the code for commercial purposes. For more information on this license and how it affects your usage of the project, please visit https://choosealicense.com/licenses/gpl-3.0/.
