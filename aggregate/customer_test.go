package aggregate_test

import (
	"ddd-go/aggregate"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Name Validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "VALID NAME",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		//Run Tests
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
