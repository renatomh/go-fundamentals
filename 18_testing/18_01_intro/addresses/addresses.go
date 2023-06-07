package addresses

import "strings"

// AddressType checks if address has a valid type and returns it
func AddressType(address string) string {
	validTypes := []string{"street", "st", "st.", "avenue",
		"av", "av.", "ave", "ave.", "road", "rd", "rd.", "way"}

	// Getting address in lower case
	lowerCaseAddress := strings.ToLower(address)
	// Splitting words (by space)
	wordsList := strings.Split(lowerCaseAddress, " ")
	// Getting he last word
	addressLastWord := wordsList[len(wordsList)-1]

	// Checking if address has a valid type
	hasValidType := false
	for _, itemType := range validTypes {
		if itemType == addressLastWord {
			hasValidType = true
		}
	}

	// If a valid type was identified
	if hasValidType {
		// Returning with the first capital letter
		return strings.Title(addressLastWord)
	}

	// Otherwise
	return "Invalid Type"
}
