package factory_pattern

type hasira struct {
	demonSlayer
}

func createHasiraDemonSlayer(name string, kc int) demonSlayerInterface {
	return &hasira{
		demonSlayer: demonSlayer{
			name: name,
			kill: kc,
		},
	}
}
