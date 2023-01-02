package addresses_test

import (
	. "intro-tests/addresses"
	"testing"
)

/*
	Creating a unit test in go.
	It'll test the address checking function
	Comparing expected results, according to pre-defined inputs
*/

// Creating the test scenario struct
type testScenario struct {
	inputAddress   string
	expectedResult string
}

// Testing function names must always start with 'Test'
// After that, we usually use the name of the function to be tested
func TestAddressType(t *testing.T) {
	// Specifying tests will run in parallel
	t.Parallel()

	// Creating the test scenarios, defining input and expected output
	testScenarios := []testScenario{
		{"ABC Street", "Street"},
		{"Hogwarts Avenue", "Avenue"},
		{"Migrants Rd.", "Rd."},
		{"Roses Plaza", "Invalid Type"},
		{"ANYWHERE AV.", "Av."},
		{"GOOFY ST.", "St."},
		{"P. Sherman 42, Wallaby Way", "Way"},
		{"", "Invalid Type"},
		{"johnsons road", "Road"},
		//{"Marksman Road", "Invalid Type"}, // This would make the test fail
	}

	// For each created scenario
	for _, scenario := range testScenarios {
		receivedResult := AddressType(scenario.inputAddress)
		// Checking if expected result is different from the returned one
		if receivedResult != scenario.expectedResult {
			// If it is, we throw a test error
			t.Errorf("For scenario \"%s\": Received data type \"%s\" is different from the expected \"%s\".",
				scenario.inputAddress,
				receivedResult,
				scenario.expectedResult,
			)
		}
	}
}

// Creating a generic simple test
func TestAny(t *testing.T) {
	// Specifying tests will run in parallel
	t.Parallel()

	// If value is true
	if 1 > 2 {
		// We throw an error and make the test fail
		t.Errorf("Test broken!")
	}
}
