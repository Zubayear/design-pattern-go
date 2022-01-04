package factory_pattern

type normal struct {
	demonSlayer
}

func createNormalDemonSlayer(name string, kc int) demonSlayerInterface {
	return &normal{
		demonSlayer: demonSlayer{
			name: name,
			kill: kc,
		},
	}
}
