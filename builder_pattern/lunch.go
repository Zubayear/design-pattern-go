package builder_pattern

type Lunch struct {
	Bread, Condiments, Dressing, Meat string
}

type LunchBuilder interface {
	Bread(bread string) LunchBuilder
	Condiments(condiments string) LunchBuilder
	Dressing(dressing string) LunchBuilder
	Meat(meat string) LunchBuilder
	Build() Lunch
}

type MyLunch struct {
	Lunch
}

func (ml *MyLunch) Bread(bread string) LunchBuilder {
	ml.Lunch.Bread = bread
	return ml
}

func (ml *MyLunch) Condiments(condiments string) LunchBuilder {
	ml.Lunch.Condiments = condiments
	return ml
}

func (ml *MyLunch) Dressing(dressing string) LunchBuilder {
	ml.Lunch.Dressing = dressing
	return ml
}

func (ml *MyLunch) Meat(meat string) LunchBuilder {
	ml.Lunch.Meat = meat
	return ml
}

func (ml *MyLunch) Build() Lunch {
	return ml.Lunch
}

func Director(lb LunchBuilder) Lunch {
	return lb.Build()
}
