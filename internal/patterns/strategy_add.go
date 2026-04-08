package patterns

import "gitlab.com/biso-ermakovkirill/practice3/internal/domain"

type AddStrategy struct{}

func (a *AddStrategy) Execute(x, y float64) (float64, error) {
	return x + y, nil
}

func (a *AddStrategy) GetName() string {
	return "add"
}

// Ensure AddStrategy implements domain.Operation
var _ domain.Operation = (*AddStrategy)(nil)