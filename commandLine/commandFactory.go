package commandLine

import "fmt"

func GetCommandToExecute(command string, cache CacheReceiver, args ...string) (Command, error) {
	switch command {
	case "GET":
		if len(args) != 1 {
			break // not enough arguments, or too many
		}

		return NewGetCommand(cache, args[0]), nil
	case "SET":
		if len(args) != 2 {
			break // not enough arguments, or too many
		}

		return NewPutCommand(cache, args[0], args[1]), nil
	case "EXIT":
		// always return the exit. do not care about the arguments
		return NewExitCommand(cache), nil
	}

	return nil, fmt.Errorf("Could not find the right command, buddy")
}
