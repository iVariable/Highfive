package commands

import (
	"google.golang.org/api/chat/v1"
)

type SorryIDontUnderstand struct {
}

func (w *SorryIDontUnderstand) Help(event *chat.DeprecatedEvent) string {
	return ""
}

func (w *SorryIDontUnderstand) Is(event *chat.DeprecatedEvent) bool {
	return event.Type == "MESSAGE"
}

func (w *SorryIDontUnderstand) Exec(event *chat.DeprecatedEvent) (*chat.Message, error) {
	return &chat.Message{
		Text: "Sorry, I did not understand. \n\nYou can get a list of available commands via \"@Highfive help\" command",
	}, nil
}
