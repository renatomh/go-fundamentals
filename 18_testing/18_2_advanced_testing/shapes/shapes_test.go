package shapes

import (
	"math"
	"testing"
)

// Function to test areas creation
func TestArea(t *testing.T) {

	// Running tests for rectangle shapes
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := rect.Area()

		if expectedArea != receivedArea {
			t.Fatalf(
				"The received area \"%f\" is different from the expected one \"%f\"",
				receivedArea,
				expectedArea,
			)
		}
	})

	// Running tests for circle shapes
	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if expectedArea != receivedArea {
			t.Fatalf(
				"The received area \"%f\" is different from the expected one \"%f\"",
				receivedArea,
				expectedArea,
			)
		}
	})
}
