package commandLine

import "fmt"

type PutCommand struct {
	Cache CacheReceiver
	Key   string
	Value string
}

func NewPutCommand(cache CacheReceiver, key string, value string) *PutCommand {
	return &PutCommand{
		Cache: cache,
		Key:   key,
		Value: value,
	}
}

func (c *PutCommand) Execute() (string, error) {
	err := c.Cache.Put(c.Key, c.Value)

	if err != nil {
		return "", fmt.Errorf("Could not perform SET operation: %v", err)
	}

	return "OK", nil
}
