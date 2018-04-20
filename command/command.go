package command

// Basic interface for type that handle a command
type Handler interface {
	Handle() error
}

// Command interface is the based typed for Command types
type Command interface {
	Handler
}

func HandleCommand(command Command) error {
	return command.Handle()
}
