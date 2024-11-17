package internal

type Client interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type ClientEntity struct {
}

func (c *ClientEntity) Get(key string) (string, error) {
	return "", nil
}

func (c *ClientEntity) Set(key, value string) error {
	return nil
}
