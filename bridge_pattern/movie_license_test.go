package bridge_pattern

import (
	"testing"
)

func TestGetMovieLicensePrice(t *testing.T) {
	type args struct {
		ml *MovieLicense
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "GetDiscountedPriceTest",
			args: args{
				ml: &MovieLicense{
					Name:        "Insomnia",
					Discount:    Military, // 4
					LicenseType: LifeLong, // 90
				},
			},
			want: 86,
		},
		{
			name: "GetDiscountedPriceTest",
			args: args{
				ml: &MovieLicense{
					Name:        "Insomnia",
					Discount:    Military, // 4
					LicenseType: Fiveday,  // 10
				},
			},
			want: 6,
		},
		{
			name: "GetDiscountedPriceTest",
			args: args{
				ml: &MovieLicense{
					Name:        "Insomnia",
					Discount:    Seinor,   // 5
					LicenseType: LifeLong, // 90
				},
			},
			want: 85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMovieLicensePrice(tt.args.ml); got != tt.want {
				t.Errorf("GetMovieLicensePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintMovieLicense(t *testing.T) {
	type args struct {
		ml *MovieLicense
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "PrintMovieLicense",
			args: args{
				ml: &MovieLicense{
					Name:        "The Dark Knight",
					Discount:    Military,
					LicenseType: LifeLong,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMovieLicense(tt.args.ml)
		})
	}
}
