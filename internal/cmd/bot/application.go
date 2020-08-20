package bot

import (
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/chat/v1"
)

type Command interface {
	Is(event *chat.DeprecatedEvent) bool
	Exec(event *chat.DeprecatedEvent) (*chat.Message, error)
	Help(event *chat.DeprecatedEvent) string
}

type Application struct {
	Log            *logrus.Entry
	Commands       []Command
	UnknownCommand Command
}

func (a *Application) ProcessEvent(event *chat.DeprecatedEvent) (*chat.Message, error) {
	for _, cmd := range a.Commands {
		if cmd.Is(event) {
			return cmd.Exec(event)
		}
	}

	if a.isHelpCommand(event) {
		return a.help(event)
	}

	return a.UnknownCommand.Exec(event)
}

func (a *Application) isHelpCommand(event *chat.DeprecatedEvent) bool {
	command := GetMessageText(event, true)

	return command == "help"
}

func (a *Application) help(event *chat.DeprecatedEvent) (*chat.Message, error) {
	helpEntries := []string{}

	for _, cmd := range a.Commands {
		if help := cmd.Help(event); help != "" {
			helpEntries = append(helpEntries, help)
		}
	}

	return &chat.Message{
		Text: "*Available commands*\n" +
			strings.Join(helpEntries, "\n\n") +
			"\n\n*Show help:* \n- help",
	}, nil
}
