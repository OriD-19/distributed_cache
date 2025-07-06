package commandLine

type Command interface {
	Execute() (string, error)
}

