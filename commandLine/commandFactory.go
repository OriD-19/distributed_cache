package commandLine

func GetCommandToExecute(command string, cache CacheReceiver, args ...string) Command {
	switch command {
	case "get":
		return NewGetCommand(cache, "")
	case "set":
		return NewSetCommand(cache, "", "")
	default:
		return nil
	}
}
