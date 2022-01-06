package singleton_pattern

import (
	"fmt"
	"sync"
)

type price interface {
	GetPrice(string) int
}

type datbase struct {
}

type foodPrice struct {
	data map[string]int
}

func (fp *foodPrice) GetPrice(item string) int {
	if len(fp.data) == 0 {
		fp.data = map[string]int{
			"Chicken": 90,
			"Meat":    190,
			"Peas":    56,
		}
	}
	return fp.data[item]
}

var db *datbase
var once sync.Once

func CalculatePrice(p price, items []string) int {
	res := 0
	for _, v := range items {
		res += p.GetPrice(v)
	}
	return res
}

func getDatabaseInstance() *datbase {
	once.Do(func() {
		fmt.Println("Creating db instance for the first time")
		db = &datbase{}
	})
	fmt.Println("Returing the previous once")
	return db
}
