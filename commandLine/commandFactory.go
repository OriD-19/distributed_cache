package commandLine

func GetCommandToExecute(command string, cache CacheReceiver, args ...string) Command {
	switch command {
	case "get":
		if len(args) != 1 {
			break // not enough arguments, or too many
		}

		return NewGetCommand(cache, args[0])
	case "set":
		if len(args) != 2 {
			break // not enough arguments, or too many
		}

		return NewPutCommand(cache, args[0], args[1])
	case "exit":
		// always return the exit. do not care about the arguments
		return NewExitCommand(cache)	
	}

	return nil
}
