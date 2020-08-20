package commands

import (
	"google.golang.org/api/chat/v1"
)

type WelcomeToTheRoom struct {
}

func (w *WelcomeToTheRoom) Help(event *chat.DeprecatedEvent) string {
	return ""
}

func (w *WelcomeToTheRoom) Is(event *chat.DeprecatedEvent) bool {
	return event.Type == "ADDED_TO_SPACE"
}

func (w *WelcomeToTheRoom) Exec(event *chat.DeprecatedEvent) (*chat.Message, error) {
	return &chat.Message{
		Text: "Hey! Thanks for adding me into the room!\n" +
			"If you want to praise someone for the job well done simply type " +
			"\"@Highfive to @person for <reason for the praise>\"\n\n" +
			"Use \"@Highfive help\" to get more details. You can also write this in the bot direct message.\n\n" +
			"Happy praising!",
	}, nil
}
