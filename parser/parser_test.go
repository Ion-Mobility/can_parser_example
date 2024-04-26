package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tc := []struct {
		Name        string
		Message     Message
		Expected    map[string]interface{}
		ExpectedErr error
	}{
		{
			Name: "valid VCS status pkt 3 message",
			Message: Message{
				BikeID:    "d5efcd9e-ebf9-4b35-adca-c80784a7da43",
				Timestamp: "2023-08-08 11:14:14.23 UTC",
				CanID:     "18edf4fa",
				Length:    5,
				Data:      "2537013596",
			},
			Expected: map[string]interface{}{
				"display_SOC":             float64(37),
				"range_km":                float64(55),
				"target_charge_hours_rem": float64(1),
				"target_charge_mins_rem":  float64(53),
				"target_charge_range_km":  float64(150),
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := Parse(tt.Message)
			assert.Equal(t, tt.Expected, got)
			assert.Equal(t, tt.ExpectedErr, err)
		})
	}
}
