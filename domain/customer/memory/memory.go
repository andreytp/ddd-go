package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (receiver *MemoryRepository) Get(uuid uuid.UUID) (aggregate.Customer, error) {

	if customerToGet, ok := receiver.customers[uuid]; ok {
		return customerToGet, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (receiver *MemoryRepository) Add(customerToAdd aggregate.Customer) error {
	if receiver.customers == nil {
		receiver.Lock()
		receiver.customers = make(map[uuid.UUID]aggregate.Customer)
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
func (receiver *MemoryRepository) Update(customerToUpdate aggregate.Customer) error {
	if _, ok := receiver.customers[customerToUpdate.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	receiver.Lock()
	receiver.customers[customerToUpdate.GetID()] = customerToUpdate
	receiver.Unlock()
	return nil
}
