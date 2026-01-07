package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeNowJakarta(t *testing.T) {
	// Call the function
	result := TimeNowJakarta()

	// 1. Basic assertion: Check if result is not empty
	assert.NotEmpty(t, result, "Result should not be empty")

	// 2. Format assertion: Check if it parses correctly with the expected format
	layout := "02 January 2006, 03:04 PM"

	// We need to load the Jakarta location to parse strictly,
	// or we can just parse it and ignore location mismatch if we only care about format structure.
	// Ideally, we load location to be precise.
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		// Fallback if local machine doesn't have zoneinfo (rare but possible in minimal docker)
		// checking length fits the format approx
		assert.Greater(t, len(result), 10)
		return
	}

	parsedTime, err := time.ParseInLocation(layout, result, loc)

	// Assert no error in parsing
	assert.NoError(t, err, "Time string should match format '02 January 2006, 03:04 PM'")
	assert.False(t, parsedTime.IsZero(), "Parsed time should not be zero")
}
