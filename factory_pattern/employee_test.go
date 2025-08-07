package factorypattern

import (
	"reflect"
	"testing"
)

func TestGetEmployee(t *testing.T) {
	type args struct {
		name   string
		age    int
		salary int
	}
	tests := []struct {
		name string
		args args
		want *Employee
	}{
		{
			name: "EmployeeFactoryTest",
			args: args{
				name:   "Syed Ibna Zubayear",
				age:    24,
				salary: 120000,
			},
			want: &Employee{
				Name:   "Syed Ibna Zubayear",
				Age:    24,
				Salary: 120000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEmployee(tt.args.name, tt.args.age, tt.args.salary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}
