package patterns

import "gitlab.com/biso-ermakovkirill/practice3/internal/domain"

type SubStrategy struct{}

func (s *SubStrategy) Execute(x, y float64) (float64, error) {
	return x - y, nil
}

func (s *SubStrategy) GetName() string {
	return "sub"
}

var _ domain.Operation = (*SubStrategy)(nil)