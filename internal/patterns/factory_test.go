package patterns

import (
	"fmt"
	"testing"

	"gitlab.com/biso-ermakovkirill/practice3/internal/domain"
)

func TestDefaultFactory(t *testing.T) {
	factory := NewDefaultFactory()

	tests := []struct {
		name     string
		opName   string
		wantType string
		wantErr  bool
	}{
		{"create add", "add", "*patterns.AddStrategy", false},
		{"create sub", "sub", "*patterns.SubStrategy", false},
		{"create mul", "mul", "*patterns.MulStrategy", false},
		{"create div", "div", "*patterns.DivStrategy", false},
		{"create pow", "pow", "*patterns.PowStrategy", false},
		{"unknown operation", "unknown", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op, err := factory.CreateOperation(tt.opName)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if op == nil {
					t.Error("expected non-nil operation")
				}
				if got := fmt.Sprintf("%T", op); got != tt.wantType {
					t.Errorf("CreateOperation(%s) type = %s, want %s", tt.opName, got, tt.wantType)
				}
			}
		})
	}
}

func TestFactoryRegisterOperation(t *testing.T) {
	factory := NewDefaultFactory()

	// Регистрируем новую операцию
	factory.RegisterOperation("test", &MockCreator{})

	op, err := factory.CreateOperation("test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if op == nil {
		t.Error("expected non-nil operation")
	}
	if op.GetName() != "mock" {
		t.Errorf("GetName() = %s, want mock", op.GetName())
	}
}

// MockCreator для тестирования регистрации
type MockCreator struct{}

func (c *MockCreator) Create() domain.Operation {
	return &MockOperation{name: "mock"}
}

type MockOperation struct {
	name string
}

func (m *MockOperation) Execute(a, b float64) (float64, error) {
	return 0, nil
}

func (m *MockOperation) GetName() string {
	return m.name
}