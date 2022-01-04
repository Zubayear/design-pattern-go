package factory_pattern

type demonSlayerInterface interface {
	setName(name string)
	killCount() int
}

type demonSlayer struct {
	name string
	kill int
}

func (d *demonSlayer) setName(name string) {
	d.name = name
}

func (d *demonSlayer) killCount() int {
	return d.kill
}
