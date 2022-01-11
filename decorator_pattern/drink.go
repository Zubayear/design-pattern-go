package decorator_pattern

import (
	"fmt"
)

type BeverageType int
type CondimentType int

const (
	DARKROAST BeverageType = iota + 1
	HOUSEBLEND
	DECAF
)

const (
	MOCHA CondimentType = iota + 1
	WHIP
	SOY
)

type Beverage interface {
	DrinkName() string
	Cost() float64
}

type CondimentDecorator interface {
	Beverage
}

type DarkRoast struct {
	name string
	cost float64
}

type HouseBlend struct {
	name string
	cost float64
}

type Decaf struct {
	name string
	cost float64
}

type Mocha struct {
	Beverage
	name string
	cost float64
}

type Soy struct {
	Beverage
	name string
	cost float64
}

type Whip struct {
	Beverage
	name string
	cost float64
}

func (s *Soy) DrinkName() string {
	return s.Beverage.DrinkName() + " with " + s.name
}

func (s *Soy) Cost() float64 {
	return s.Beverage.Cost() + s.cost
}

func (w *Whip) DrinkName() string {
	return w.Beverage.DrinkName() + " with " + w.name
}

func (w *Whip) Cost() float64 {
	return w.Beverage.Cost() + w.cost
}

func (dc *Decaf) DrinkName() string {
	return dc.name
}

func (dc *Decaf) Cost() float64 {
	return dc.cost
}

func (hb *HouseBlend) DrinkName() string {
	return hb.name
}

func (hb *HouseBlend) Cost() float64 {
	return hb.cost
}

func (m *Mocha) DrinkName() string {
	return m.Beverage.DrinkName() + " with " + m.name
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + m.cost
}

func (dr *DarkRoast) DrinkName() string {
	return dr.name
}

func (dr *DarkRoast) Cost() float64 {
	return dr.cost
}

func NewDarkRoast(name string, cost float64) *DarkRoast {
	return &DarkRoast{
		name: name,
		cost: cost,
	}
}

func NewHouseBlend(name string, cost float64) *HouseBlend {
	return &HouseBlend{
		name: name,
		cost: cost,
	}
}

func NewDecaf(name string, cost float64) *Decaf {
	return &Decaf{
		name: name,
		cost: cost,
	}
}

func NewBeverage(name string, cost float64, bt BeverageType) Beverage {
	switch bt {
	case DARKROAST:
		return NewDarkRoast(name, cost)
	case HOUSEBLEND:
		return NewHouseBlend(name, cost)
	case DECAF:
		return NewDecaf(name, cost)
	default:
		return nil
	}
}

func NewMocha(b Beverage, name string, cost float64) *Mocha {
	return &Mocha{
		Beverage: b,
		name:     name,
		cost:     cost,
	}
}

func NewSoy(b Beverage, name string, cost float64) *Soy {
	return &Soy{
		Beverage: b,
		name:     name,
		cost:     cost,
	}
}

func NewWhip(b Beverage, name string, cost float64) *Whip {
	return &Whip{
		Beverage: b,
		name:     name,
		cost:     cost,
	}
}

func NewCondiment(b Beverage, name string, cost float64, bt CondimentType) Beverage {
	switch bt {
	case MOCHA:
		return NewMocha(b, name, cost)
	case SOY:
		return NewSoy(b, name, cost)
	case WHIP:
		return NewWhip(b, name, cost)
	default:
		return nil
	}
}

func GetCost(b Beverage) string {
	cost := b.Cost()
	return fmt.Sprintf("%.2f", cost)
}

func GetName(b Beverage) string {
	return b.DrinkName()
}
