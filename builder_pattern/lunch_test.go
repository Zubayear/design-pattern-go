package builder_pattern

import (
	"reflect"
	"testing"
)

func TestLunchBuilder_BuildLunch(t *testing.T) {
	type fields struct {
		bread      string
		condiments string
		dressing   string
		meat       string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Lunch
	}{
		{
			name: "LunchTest",
			fields: fields{
				bread:      "Whole Wheat Bread",
				condiments: "Lettuce",
				dressing:   "Mayo",
				meat:       "beef",
			},
			want: &Lunch{
				bread:      "Whole Wheat Bread",
				condiments: "Lettuce",
				dressing:   "Mayo",
				meat:       "beef",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &LunchBuilder{
				bread:      tt.fields.bread,
				condiments: tt.fields.condiments,
				dressing:   tt.fields.dressing,
				meat:       tt.fields.meat,
			}
			if got := lb.BuildLunch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LunchBuilder.BuildLunch() = %v, want %v", got, tt.want)
			}
		})
	}
}
