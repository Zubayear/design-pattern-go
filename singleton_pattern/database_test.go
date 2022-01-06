package singleton_pattern

import (
	"reflect"
	"testing"
)

func Test_getDatabaseInstance(t *testing.T) {
	tests := []struct {
		name string
		want *datbase
	}{
		// TODO: Add test cases.
		{
			name: "SingletonTest",
			want: &datbase{},
		},
		{
			name: "SingletonTest",
			want: &datbase{},
		},
		{
			name: "SingletonTest",
			want: &datbase{},
		},
		{
			name: "SingletonTest",
			want: &datbase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDatabaseInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDatabaseInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePrice(t *testing.T) {
	type args struct {
		p     price
		items []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "SingleTonWithDIP",
			args: args{
				p:     &foodPrice{},
				items: []string{"Chicken", "Peas"},
			},
			want: 90 + 56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePrice(tt.args.p, tt.args.items); got != tt.want {
				t.Errorf("CalculatePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
