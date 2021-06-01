// !!! This file forbidden for updates during solving task !!!

package volume_calculator

import (
	"github.com/stretchr/testify/assert"
	"volume_calculator/example/calculator"

	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name           string
		shape          calculator.VolumeCalculator
		expectedVolume float64
	}{
		{
			name:           "Cube volume",
			shape:          calculator.NewCube(3),
			expectedVolume: 27,
		},
		{
			name:           "Pyramid volume",
			shape:          calculator.NewPyramid(1, 2, 3),
			expectedVolume: 2,
		},
		{
			name:           "Sphere volume",
			shape:          calculator.NewSphere(2),
			expectedVolume: 33.51,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actualVolume := calculator.CalculateVolume(tt.shape)
			assert.Equal(t, tt.expectedVolume, actualVolume)
		})
	}

}
