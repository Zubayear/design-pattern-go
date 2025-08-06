package adapterpattern

import "fmt"

type duck interface {
	fly()
	quack()
}

type turkey interface {
	gobble()
	fly()
}

type mallardDuck struct{}
type wildTurkey struct{}

// Adapt to the duck interface
// Compose with turkey interface
// Call the underlying methods
type turkeyAdapter struct {
	turkey
}

func (md *mallardDuck) fly() {
	fmt.Println("Mallard Duck is flying")
}

func (md *mallardDuck) quack() {
	fmt.Println("Mallard Duck is quacking")
}

func (wt *wildTurkey) fly() {
	fmt.Println("Wild Turkey is flying short distance")
}

func (wt *wildTurkey) gobble() {
	fmt.Println("Mallard Duck is gobbling")
}

func (ta *turkeyAdapter) fly() {
	ta.turkey.fly()
}

func (ta *turkeyAdapter) quack() {
	ta.turkey.gobble()
}
