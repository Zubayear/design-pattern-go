package factory_pattern

import (
	"reflect"
	"testing"
)

func Test_newDemonSlayer(t *testing.T) {
	type args struct {
		st   SlayerType
		name string
		kc   int
	}
	tests := []struct {
		name    string
		args    args
		want    demonSlayerInterface
		wantErr bool
	}{
		{
			name: "FactoryTest",
			args: args{
				st:   2,
				name: "Tanjiro",
				kc:   10,
			},
			want: &normal{
				demonSlayer: demonSlayer{
					name: "Tanjiro",
					kill: 10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDemonSlayer(tt.args.st, tt.args.name, tt.args.kc)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDemonSlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDemonSlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
