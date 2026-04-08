package patterns

import (
	"fmt"
	"github.com/Kiryadx/practice3_Ermakov/internal/domain"
)

// OperationCreator — интерфейс создателя стратегии
type OperationCreator interface {
	Create() domain.Operation
}

// OperationFactory — интерфейс фабрики
type OperationFactory interface {
	CreateOperation(name string) (domain.Operation, error)
	RegisterOperation(name string, creator OperationCreator)
}

// DefaultFactory — конкретная фабрика
type DefaultFactory struct {
	creators map[string]OperationCreator
}

// NewDefaultFactory создаёт фабрику с предустановленными операциями
func NewDefaultFactory() *DefaultFactory {
	f := &DefaultFactory{
		creators: make(map[string]OperationCreator),
	}
	f.registerDefaultOperations()
	return f
}

func (f *DefaultFactory) registerDefaultOperations() {
	f.RegisterOperation("add", &AddCreator{})
	f.RegisterOperation("sub", &SubCreator{})
	f.RegisterOperation("mul", &MulCreator{})
	f.RegisterOperation("div", &DivCreator{})
	f.RegisterOperation("pow", &PowCreator{})
}

func (f *DefaultFactory) RegisterOperation(name string, creator OperationCreator) {
	f.creators[name] = creator
}

func (f *DefaultFactory) CreateOperation(name string) (domain.Operation, error) {
	creator, exists := f.creators[name]
	if !exists {
		return nil, fmt.Errorf("неизвестная операция: %s", name)
	}
	return creator.Create(), nil
}

// Конкретные создатели
type AddCreator struct{}
func (c *AddCreator) Create() domain.Operation { return &AddStrategy{} }

type SubCreator struct{}
func (c *SubCreator) Create() domain.Operation { return &SubStrategy{} }

type MulCreator struct{}
func (c *MulCreator) Create() domain.Operation { return &MulStrategy{} }

type DivCreator struct{}
func (c *DivCreator) Create() domain.Operation { return &DivStrategy{} }

type PowCreator struct{}
func (c *PowCreator) Create() domain.Operation { return &PowStrategy{} }