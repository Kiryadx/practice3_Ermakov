package patterns

import (
	"testing"
)

func TestAddStrategy(t *testing.T) {
	strategy := &AddStrategy{}

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
		{"mixed", -5, 3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strategy.Execute(tt.a, tt.b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Execute(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}

	if strategy.GetName() != "add" {
		t.Errorf("GetName() = %s, want add", strategy.GetName())
	}
}

func TestSubStrategy(t *testing.T) {
	strategy := &SubStrategy{}

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive", 10, 3, 7},
		{"negative", -5, -2, -3},
		{"zero", 0, 5, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strategy.Execute(tt.a, tt.b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Execute(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivStrategy(t *testing.T) {
	strategy := &DivStrategy{}

	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"normal", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strategy.Execute(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Execute(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestPowStrategy(t *testing.T) {
	strategy := &PowStrategy{}

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"2^3", 2, 3, 8},
		{"5^0", 5, 0, 1},
		{"0^5", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strategy.Execute(tt.a, tt.b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Execute(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}