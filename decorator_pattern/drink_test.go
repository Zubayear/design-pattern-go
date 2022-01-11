package decorator_pattern

import (
	"testing"
)

func TestGetName(t *testing.T) {
	type args struct {
		b Beverage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "BeverageNameTest",
			args: args{
				b: NewCondiment(NewBeverage("Dark Roast", 5.9, BeverageType(MOCHA)), "Mocha", 7.8, CondimentType(DARKROAST)),
			},
			want: "Dark Roast with Mocha",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetName(tt.args.b); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCost(t *testing.T) {
	type args struct {
		b Beverage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "BeverageCostTest",
			args: args{
				b: NewCondiment(NewBeverage("Dark Roast", 5.9, BeverageType(MOCHA)), "Mocha", 7.8, CondimentType(DARKROAST)),
			},
			want: "13.70",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCost(tt.args.b); got != tt.want {
				t.Errorf("GetCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
