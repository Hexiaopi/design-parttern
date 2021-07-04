package factory_method

import "testing"

func TestOperator(t *testing.T) {
	tests := []struct {
		name string
		OperatorFactory
		a    int
		b    int
		want int
	}{
		{name: "加法二元运算符", OperatorFactory: PlusOperatorFactory{}, a: 1, b: 2, want: 3},
		{name: "减法二元运算符", OperatorFactory: MinusOperatorFactory{}, a: 3, b: 2, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compute(tt.OperatorFactory, tt.a, tt.b)
			if got != tt.want {
				t.Errorf("computer a:%d and b:%d got:%d,want:%d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
