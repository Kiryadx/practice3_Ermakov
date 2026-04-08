package patterns

import (
	"math"
	"gitlab.com/biso-ermakovkirill/practice3/internal/domain"
)

type PowStrategy struct{}

func (p *PowStrategy) Execute(x, y float64) (float64, error) {
	return math.Pow(x, y), nil
}

func (p *PowStrategy) GetName() string {
	return "pow"
}

var _ domain.Operation = (*PowStrategy)(nil)