package patterns

import (
	"fmt"
	"gitlab.com/biso-ermakovkirill/practice3/internal/domain"
)

type DivStrategy struct{}

func (d *DivStrategy) Execute(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return x / y, nil
}

func (d *DivStrategy) GetName() string {
	return "div"
}

var _ domain.Operation = (*DivStrategy)(nil)