package patterns

import "gitlab.com/biso-ermakovkirill/practice3/internal/domain"

type MulStrategy struct{}

func (m *MulStrategy) Execute(x, y float64) (float64, error) {
	return x * y, nil
}

func (m *MulStrategy) GetName() string {
	return "mul"
}

var _ domain.Operation = (*MulStrategy)(nil)