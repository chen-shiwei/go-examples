package option

type Person struct {
	options
}

func NewPerson(ops ...Option) *Person {
	options := options{}
	for _, o := range ops {
		o.apply(&options)
	}
	p := Person{options: options}
	return &p
}
