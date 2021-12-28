package aggregate

import "testing"

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: ErrMissingValues,
		},
		{
			test:        "validvalues",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		product, err := NewProduct(tc.name, tc.description, tc.price)
		if err != tc.expectedErr {
			t.Errorf("Expected error was %v but get %v", tc.expectedErr, err)
			continue
		}

		if tc.expectedErr != nil {
			continue
		}

		if product.item.Name != tc.name || product.item.Description != tc.description || product.price != tc.price {
			t.Errorf("Expected error was %v but get %v", tc.expectedErr, ErrWrongValues)
		}

	}
}
