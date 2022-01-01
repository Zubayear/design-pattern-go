package builder_pattern

type Lunch struct {
	bread, condiments, dressing, meat string
}

// Builder will have methods attach to set the fields and return the builder for fluent builder
// Also, will have build to return the ultimate product
type LunchBuilder struct {
	bread, condiments, dressing, meat string
}

func NewLunchBuilder() *LunchBuilder {
	return &LunchBuilder{}
}

func (lb *LunchBuilder) Bread(bread string) *LunchBuilder {
	lb.bread = bread
	return lb
}

func (lb *LunchBuilder) Condiments(condiments string) *LunchBuilder {
	lb.condiments = condiments
	return lb
}

func (lb *LunchBuilder) Dressing(dressing string) *LunchBuilder {
	lb.dressing = dressing
	return lb
}

func (lb *LunchBuilder) Meat(meat string) *LunchBuilder {
	lb.meat = meat
	return lb
}

func (lb *LunchBuilder) BuildLunch() *Lunch {
	return &Lunch{
		bread:      lb.bread,
		condiments: lb.condiments,
		dressing:   lb.dressing,
		meat:       lb.meat,
	}
}
