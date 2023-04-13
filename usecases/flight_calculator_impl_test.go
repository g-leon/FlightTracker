package usecases

import (
	"FlightTracker/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name         string
		inputFlights []*entities.Flight
		expected     *entities.Flight
	}{
		{
			name: "single flight",
			inputFlights: []*entities.Flight{
				entities.NewFlight([]string{"A", "B"}),
			},
			expected: entities.NewFlight([]string{"A", "B"}),
		},
		{
			name: "multiple flights with clear start and end",
			inputFlights: []*entities.Flight{
				entities.NewFlight([]string{"A", "B"}),
				entities.NewFlight([]string{"B", "C"}),
			},
			expected: entities.NewFlight([]string{"A", "C"}),
		},
		{
			name: "multiple flights with circular route",
			inputFlights: []*entities.Flight{
				entities.NewFlight([]string{"A", "B"}),
				entities.NewFlight([]string{"B", "C"}),
				entities.NewFlight([]string{"C", "A"}),
			},
			expected: entities.NewFlight([]string{"", ""}),
		},
		{
			name:         "empty input",
			inputFlights: []*entities.Flight{},
			expected:     entities.NewFlight([]string{"", ""}),
		},
	}

	fc := &FlightCalculatorImpl{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := fc.Calculate(tc.inputFlights)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Airports(), result.Airports())
		})
	}
}
