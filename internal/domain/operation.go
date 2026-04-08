package domain

// Operation — интерфейс стратегии (паттерн Strategy)
type Operation interface {
	Execute(a, b float64) (float64, error)
	GetName() string
}