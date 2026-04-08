package calculator

import (
	"gitlab.com/biso-ermakovkirill/practice3/internal/patterns"
)

// Calculator — контекст, использующий стратегии
type Calculator struct {
	factory       patterns.OperationFactory
	enableLogging bool
}

// NewCalculator создаёт новый калькулятор
func NewCalculator(enableLogging bool) *Calculator {
	return &Calculator{
		factory:       patterns.NewDefaultFactory(),
		enableLogging: enableLogging,
	}
}

// Calculate выполняет операцию с заданным именем
func (c *Calculator) Calculate(opName string, a, b float64) (float64, error) {
	// Фабричный метод создаёт стратегию
	op, err := c.factory.CreateOperation(opName)
	if err != nil {
		return 0, err
	}

	// Если логирование включено — оборачиваем стратегию в декоратор
	if c.enableLogging {
		op = patterns.NewLoggingDecorator(op)
	}

	// Выполняем операцию
	return op.Execute(a, b)
}

// SetEnableLogging включает/выключает логирование
func (c *Calculator) SetEnableLogging(enabled bool) {
	c.enableLogging = enabled
}

// GetAvailableOperations возвращает список доступных операций
func (c *Calculator) GetAvailableOperations() []string {
	return []string{"add", "sub", "mul", "div", "pow"}
}