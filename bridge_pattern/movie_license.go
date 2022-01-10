package bridge_pattern

import "fmt"

type LicenseType int
type Discount int

const (
	Fiveday LicenseType = iota + 1
	LifeLong
)

const (
	Military Discount = iota + 1
	Seinor
)

type MovieLicense struct {
	Name string
	Discount
	LicenseType
}

func (ml *MovieLicense) GetDiscount() int {
	switch ml.Discount {
	case Seinor:
		return 5
	case Military:
		return 4
	default:
		return 0
	}
}

func (ml *MovieLicense) GetBasePrice() int {
	switch ml.LicenseType {
	case Fiveday:
		return 10
	case LifeLong:
		return 90
	default:
		return 3
	}
}

func (ml *MovieLicense) GetPrice() int {
	return ml.GetBasePrice() - ml.GetDiscount()
}

func NewMovieLicense(name string, discount Discount, licenseType LicenseType) *MovieLicense {
	return &MovieLicense{
		Name:        name,
		Discount:    discount,
		LicenseType: licenseType,
	}
}

func PrintMovieLicense(ml *MovieLicense) {
	fmt.Println(ml)
}

func GetMovieLicensePrice(ml *MovieLicense) int {
	return ml.GetPrice()
}
