package option

type options struct {
	Name string
	Age  int
}

type Option interface {
	apply(c *options)
}

type ConfigFunc func(c *options)

func (f ConfigFunc) apply(c *options) {
	f(c)
}

func WithName(name string) Option {
	return ConfigFunc(func(c *options) {
		c.Name = name
	})
}

func WithAge(age int) Option {
	return ConfigFunc(func(c *options) {
		c.Age = age
	})
}
