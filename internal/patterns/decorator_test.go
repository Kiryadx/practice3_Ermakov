package patterns

import (
	"testing"
)

func TestLoggingDecorator(t *testing.T) {
	// Создаём реальную стратегию
	strategy := &AddStrategy{}

	// Оборачиваем в декоратор
	decorator := NewLoggingDecorator(strategy)

	// Проверяем, что декоратор реализует интерфейс
	if decorator.GetName() != "add" {
		t.Errorf("GetName() = %s, want add", decorator.GetName())
	}

	// Проверяем выполнение
	result, err := decorator.Execute(2, 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Execute() = %f, want 5", result)
	}
}

func TestLoggingDecoratorWithError(t *testing.T) {
	// Создаём стратегию, которая возвращает ошибку
	strategy := &DivStrategy{}
	decorator := NewLoggingDecorator(strategy)

	// Деление на ноль — должно вызвать ошибку
	_, err := decorator.Execute(10, 0)
	if err == nil {
		t.Error("expected error, got nil")
	}
}