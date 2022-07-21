package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// table tests
func TestDo(t *testing.T) {
	tests := map[string]struct {
		char           string
		number         int
		boolean        bool
		expectedResult string
		expectedErr    string
	}{
		"success": {
			char:           "a",
			number:         3,
			boolean:        true,
			expectedResult: "[3A]",
		},
		"successMinLength": {
			char:           "a",
			number:         21,
			boolean:        true,
			expectedResult: "A",
		},
		"invalidChar": {
			char:           "e",
			number:         3,
			boolean:        true,
			expectedResult: "",
			expectedErr:    "invalid s",
		},
		"invalidNumber": {
			char:        "a",
			number:      7,
			boolean:     true,
			expectedErr: "invalid i",
		},
		"booleanFalse": {
			char:           "a",
			number:         21,
			boolean:        false,
			expectedResult: "a",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualResult, err := Do(tc.char, tc.number, tc.boolean)
			assert.Equal(t, tc.expectedResult, actualResult)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// primitive approach, let it be here :)
//func TestDoSuccess(t *testing.T) {
//	actualResult, err := Do("a", 3, true)
//	expectedResult := "[3A]"
//
//	assert.Equal(t, expectedResult, actualResult)
//	assert.NoError(t, err)
//}
//
//func TestDoInvalidS(t *testing.T) {
//	actualResult, err := Do("abc  ", 3, true)
//	expectedResult := ""
//
//	assert.Equal(t, expectedResult, actualResult)
//	assert.EqualError(t, err, "invalid s")
//}
