package commands

import (
	"strings"

	"google.golang.org/api/chat/v1"
)

type MeCommand struct {
}

func (w *MeCommand) Help(event *chat.DeprecatedEvent) string {
	return "Good old IRC /me command"
}

func (w *MeCommand) Is(event *chat.DeprecatedEvent) bool {
	return event.Type == "MESSAGE" && strings.HasPrefix(event.Message.Text, "/me ")
}

func (w *MeCommand) Exec(event *chat.DeprecatedEvent) (*chat.Message, error) {
	return &chat.Message{
		Text: "*"+event.User.DisplayName + "* " + event.Message.Text[4:],
	}, nil
}
