package factory_pattern

import "errors"

type SlayerType int

const (
	Hasira SlayerType = iota + 1
	Normal
)

func newDemonSlayer(st SlayerType, name string, kc int) (demonSlayerInterface, error) {
	switch st {
	case Hasira:
		return createHasiraDemonSlayer(name, kc), nil
	case Normal:
		return createNormalDemonSlayer(name, kc), nil
	default:
		return nil, errors.New("invalid type of demon slayer")
	}
}
