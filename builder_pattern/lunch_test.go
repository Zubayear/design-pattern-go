package builderpattern

import (
	"reflect"
	"testing"
)

func TestDirector(t *testing.T) {
	type args struct {
		lb LunchBuilder
	}
	tests := []struct {
		name string
		args args
		want Lunch
	}{
		{
			name: "Lunch",
			args: args{
				lb: &MyLunch{
					Lunch: Lunch{
						Bread:      "Brioche",
						Condiments: "Mustard",
						Dressing:   "Mayonnaise",
						Meat:       "Beef",
					},
				},
			},
			want: Lunch{
				Bread:      "Brioche",
				Condiments: "Mustard",
				Dressing:   "Mayonnaise",
				Meat:       "Beef",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Director(tt.args.lb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Director() = %v, want %v", got, tt.want)
			}
		})
	}
}
