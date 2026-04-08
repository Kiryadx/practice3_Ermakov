package calculator

import (
	"testing"
)

func TestCalculatorCalculate(t *testing.T) {
	calc := NewCalculator(false)

	tests := []struct {
		name    string
		op      string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"add", "add", 10, 5, 15, false},
		{"sub", "sub", 10, 5, 5, false},
		{"mul", "mul", 10, 5, 50, false},
		{"div", "div", 10, 5, 2, false},
		{"pow", "pow", 2, 3, 8, false},
		{"division by zero", "div", 10, 0, 0, true},
		{"unknown operation", "unknown", 10, 5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Calculate(tt.op, tt.a, tt.b)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculate(%s, %f, %f) = %f, want %f", tt.op, tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculatorWithLogging(t *testing.T) {
	calc := NewCalculator(true)

	result, err := calc.Calculate("add", 10, 5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 15 {
		t.Errorf("result = %f, want 15", result)
	}
}

func TestCalculatorGetAvailableOperations(t *testing.T) {
	calc := NewCalculator(false)

	ops := calc.GetAvailableOperations()

	expectedOps := []string{"add", "sub", "mul", "div", "pow"}

	if len(ops) != len(expectedOps) {
		t.Errorf("len(ops) = %d, want %d", len(ops), len(expectedOps))
	}

	for i, op := range ops {
		if op != expectedOps[i] {
			t.Errorf("ops[%d] = %s, want %s", i, op, expectedOps[i])
		}
	}
}