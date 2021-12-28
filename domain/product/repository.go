package product

import (
	"ddd-go/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("the product not found")
	ErrProductAlreadyExist = errors.New("the product already exist")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
