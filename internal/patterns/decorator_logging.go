package patterns

import (
	"fmt"
	"github.com/Kiryadx/practice3_Ermakov/internal/domain"
)

// LoggingDecorator — паттерн Decorator для добавления логирования
type LoggingDecorator struct {
	wrapped domain.Operation
}

// NewLoggingDecorator создаёт обёртку с логированием
func NewLoggingDecorator(op domain.Operation) *LoggingDecorator {
	return &LoggingDecorator{wrapped: op}
}

func (l *LoggingDecorator) Execute(a, b float64) (float64, error) {
	fmt.Printf("[LOG] Вызов: %s(%.2f, %.2f)\n", l.wrapped.GetName(), a, b)

	result, err := l.wrapped.Execute(a, b)

	if err != nil {
		fmt.Printf("[LOG] Ошибка: %v\n", err)
	} else {
		fmt.Printf("[LOG] Результат: %.2f\n", result)
	}

	return result, err
}

func (l *LoggingDecorator) GetName() string {
	return l.wrapped.GetName()
}

// Ensure LoggingDecorator implements domain.Operation
var _ domain.Operation = (*LoggingDecorator)(nil)