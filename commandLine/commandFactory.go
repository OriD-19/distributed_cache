package commandLine

import "fmt"

func GetCommandToExecute(command string, cache CacheReceiver, args ...string) (Command, error) {
	switch command {
	case "GET":
		if len(args) != 1 {
			return nil, fmt.Errorf("Syntax Error: Wrong number of arguments for GET command (expected 1, got %d)", len(args))
		}

		return NewGetCommand(cache, args[0]), nil
	case "SET":
		if len(args) != 2 {
			return nil, fmt.Errorf("Syntax Error: Wrong number of arguments for SET command (expected 2, got %d)", len(args))
		}

		return NewPutCommand(cache, args[0], args[1]), nil
	case "EXIT":
		// always return the exit. do not care about the arguments
		return NewExitCommand(cache), nil
	}

	return nil, fmt.Errorf("Syntax Error: Unkown command")
}
