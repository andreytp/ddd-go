package memory

import (
	"ddd-go/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (receiver *MemoryRepository) Get(uuid uuid.UUID) (customer.Customer, error) {

	if customerToGet, ok := receiver.customers[uuid]; ok {
		return customerToGet, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (receiver *MemoryRepository) Add(customerToAdd customer.Customer) error {
	if receiver.customers == nil {
		receiver.Lock()
		receiver.customers = make(map[uuid.UUID]customer.Customer)
		receiver.Unlock()
	}
	if _, ok := receiver.customers[customerToAdd.GetID()]; ok {
		return fmt.Errorf("customer alredy exists: %w", customer.ErrFailedToAddCustomer)
	}
	receiver.Lock()
	receiver.customers[customerToAdd.GetID()] = customerToAdd
	receiver.Unlock()
	return nil
}
func (receiver *MemoryRepository) Update(customerToUpdate customer.Customer) error {
	if _, ok := receiver.customers[customerToUpdate.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	receiver.Lock()
	receiver.customers[customerToUpdate.GetID()] = customerToUpdate
	receiver.Unlock()
	return nil
}
