package container

func NewContainer() *Container {
	c := Container{
		Entries: make(map[string]interface{}),
	}
	return &c
}

type Container struct {
	Entries map[string]interface{}
}

func (c *Container) Get(key string) interface{} {
	v, ok := c.Entries[key]
	if !ok {
		return nil
	}
	return v
}

func (c *Container) Set(key string, value interface{}) {
	c.Entries[key] = value
}
