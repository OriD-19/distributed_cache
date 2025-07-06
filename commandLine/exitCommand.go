package commandLine

type ExitCommand struct {
	Cache CacheReceiver
}

func NewExitCommand(cache CacheReceiver) *ExitCommand {
	return &ExitCommand{
		Cache: cache,
	}
}

func (c *ExitCommand) Execute() (string, error) {
	return "BYE", nil
}
