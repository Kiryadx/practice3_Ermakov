package domain

import "testing"

func TestOperationInterface(t *testing.T) {
	// Проверяем, что интерфейс работает с моком
	mock := &MockOperation{
		Name:   "test",
		Result: 42,
		Err:    nil,
	}

	var op Operation = mock

	if op.GetName() != "test" {
		t.Errorf("GetName() = %s, want test", op.GetName())
	}

	result, err := op.Execute(1, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("Execute() = %f, want 42", result)
	}
}

type MockOperation struct {
	Name   string
	Result float64
	Err    error
}

func (m *MockOperation) Execute(a, b float64) (float64, error) {
	return m.Result, m.Err
}

func (m *MockOperation) GetName() string {
	return m.Name
}