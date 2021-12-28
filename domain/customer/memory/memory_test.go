package memory

import (
	"ddd-go/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemoryRepository_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	newCustomer, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := newCustomer.GetID()
	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: newCustomer,
		},
	}
	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

func TestMemoryRepository_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		addCustomer customer.Customer
		expectedErr error
	}
	newCustomer, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []testCase{
		{
			name:        "Add customer",
			addCustomer: newCustomer,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]customer.Customer{},
			}

			err = repo.Add(tc.addCustomer)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(newCustomer.GetID())
			if err != nil {
				t.Fatal(err)
			}

			if found.GetID() != newCustomer.GetID() {
				t.Errorf("Expected %v, got %v", newCustomer.GetID(), found.GetID())
			}
		})

	}
}
