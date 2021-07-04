package simple_factory

import "testing"

func TestNewGun(t *testing.T) {
	type wants struct {
		name  string
		power int
	}
	tests := []struct {
		name  string
		args  string
		wants wants
	}{
		{name: "ak47-test", args: "ak47", wants: wants{
			name:  "AK47 Gun",
			power: 4,
		}},
		{name: "m4a1-test", args: "m4a1", wants: wants{
			name:  "M4A1 Gun",
			power: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gun := NewGun(tt.args)
			name := gun.getName()
			power := gun.getPower()
			if name != tt.wants.name || power != tt.wants.power {
				t.Errorf("NewGun name = %s and power = %d, want %v", name, power, tt.wants)
			}
		})
	}
}
