package repository

type config struct {
	isMock bool
}

type option func(*config)

func OptionIsMock() option {
	return func(c *config) {
		c.isMock = true
	}
}
