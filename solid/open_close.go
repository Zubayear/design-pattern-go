package solid

import (
	"strings"
	"time"
)

// without Open-close principle we would attach new methods to existing type
//type Duck struct {
//
//}
//
//func (d *Duck) MigrationFly() {
//	// ....
//}
//
//func (d *Duck) InjuredFly() {
//
//}

type FlyBehaviour interface {
	Fly() string
}

// Duck compose FlyBehaviour so that we can provide implementation
// so, when new implementation comes then we just extend i.e. implement the interface
// had we not compose Duck with FlyBehaviour interface then we would need to attach 
// new methods like InjuredDuckFly to Duck
// which violates type open for extension but close for modification
type Duck struct {
	FlyBehaviour
}

func NewDuck(flyBehaviour FlyBehaviour) *Duck {
	return &Duck{FlyBehaviour: flyBehaviour}
}

func DuckTest(duck Duck) string {
	return duck.Fly()
}

type MigrationFly struct {
}

func (m *MigrationFly) Fly() string {
	sb := strings.Builder{}
	sb.WriteString("Migration Fly @")
	sb.WriteString(time.Now().Format(time.ANSIC))
	return sb.String()
}

// InjuredFly new type which extend
// i.e. open for extension but close for modification
type InjuredFly struct {
}

func (i *InjuredFly) Fly() string {
	sb := strings.Builder{}
	sb.WriteString("Injured Fly @")
	sb.WriteString(time.Now().Format(time.ANSIC))
	return sb.String()
}

// The Open-Closed Principle (OCP) is a software design principle that states
// that software entities (classes, modules, functions, etc.) should be open for
// extension but closed for modification. This means that existing code should not
// be modified in order to add new functionality or behavior, but instead should be
// extended through new code that builds on top of the existing code.

// Suppose we have a `Payment` interface that represents a payment method,
// and we want to be able to process payments through different payment gateways.
// We could define a `ProcessPayment` function that takes a `Payment` as input and returns a Transaction:
type Payment interface {
	Process(amount int) (*Transaction, error)
}

type Transaction struct {
	ID     int
	Amount int
}

func ProcessPayment(p Payment, amount int) (*Transaction, error) {
	return p.Process(amount)
}

// We could then define several different types of payment gateways that
// implement the Payment interface, such as PayPal, Stripe, and Authorize.Net:
type PayPal struct {
	Username string
	Password string
}

func (p *PayPal) Process(amount int) (*Transaction, error) {
	return &Transaction{ID: 1, Amount: amount}, nil
}

type Stripe struct {
	APIKey string
}

func (s *Stripe) Process(amount int) (*Transaction, error) {
	// Code to process payment through Stripe
	return &Transaction{ID: 1, Amount: amount}, nil
}

type AuthorizeNet struct {
	APIKey string
}

func (a *AuthorizeNet) Process(amount int) (*Transaction, error) {
	// Code to process payment through Authorize.Net
	return &Transaction{ID: 1, Amount: amount}, nil
}

// Now, suppose we want to add a new payment gateway, Braintree, and
// process a payment through it. Instead of modifying the existing
// ProcessPayment function or the Payment interface, we can create a new
// Braintree type that implements the Payment interface and define its Process method:
type Braintree struct {
	MerchantID string
	PublicKey  string
	PrivateKey string
}

func (bt *Braintree) Process(amount int) (*Transaction, error) {
	return &Transaction{ID: 1, Amount: amount}, nil
}

// We can then use the existing ProcessPayment function to process a
// payment through Braintree without modifying the existing code:

// braintree := Braintree{
//     MerchantID: "12345",
//     PublicKey:  "abcde",
//     PrivateKey: "fghij",
// }
// transaction, err := ProcessPayment(braintree, 100.0)

// By designing our code to be open for extension but closed for modification,
// we have made our code more flexible and easier to maintain. We can add
// new payment gateways and process payments through them without
// modifying existing code, which reduces the risk of introducing bugs or breaking existing functionality.
