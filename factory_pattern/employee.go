package factory_pattern

type Employee struct {
	Name   string
	Age    int
	Salary int
}

func NewEmployeeFactory(age, salary int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:   name,
			Age:    age,
			Salary: salary,
		}
	}
}

func GetEmployee(name string, age, salary int) *Employee {
	return NewEmployeeFactory(age, salary)(name)
}
