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

// Duck compose FlyBehaviour so that
// we can provide implementation
// so, when new implementation comes then
// we just extend i.e. implement the interface
// had we not compose Duck with FlyBehaviour interface
// then we would need to attach new methods like InjuredDuckFly to Duck
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
