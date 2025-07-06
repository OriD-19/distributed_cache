package commandLine

import "fmt"

type Command interface {
	execute()
}

type PutCommand struct {
	Cache CacheReceiver
	Key   string
	Value string
}

func NewSetCommand(cache CacheReceiver, key string, value string) *PutCommand {
	return &PutCommand{
		Cache: cache,
		Key:   key,
		Value: value,
	}
}

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

func (c *PutCommand) execute() error {
	err := c.Cache.Put(c.Key, c.Value)

	if err != nil {
		return fmt.Errorf("Could not perform SET operation: %v", err)
	}

	return nil
}

func (c *GetCommand) execute() (string, error) {
	value, err := c.Cache.Get(c.Key)

	if err != nil {
		return "", fmt.Errorf("Could not perform GET operation: %v", err)
	}

	return value, nil
}
