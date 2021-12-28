package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range products {
		products = append(products, product)
	}
	return products, nil
}

func (m MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := m.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (m MemoryProductRepository) Add(newProduct aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}
	m.products[newProduct.GetID()] = newProduct

	return nil
}

func (m MemoryProductRepository) Update(upProduct aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[upProduct.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.products[upProduct.GetID()] = upProduct
	return nil
}

func (m MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(m.products, id)
	return nil
}
