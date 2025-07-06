package commandLine

import "fmt"

type GetCommand struct {
	Cache CacheReceiver
	Key   string
}

func NewGetCommand(cache CacheReceiver, key string) *GetCommand {
	return &GetCommand{
		Cache: cache,
		Key:   key,
	}
}

func (c *GetCommand) Execute() (string, error) {
	value, err := c.Cache.Get(c.Key)

	if err != nil {
		return "", fmt.Errorf("Could not perform GET operation: %v", err)
	}

	return value, nil
}
