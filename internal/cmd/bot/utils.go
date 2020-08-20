package bot

import (
	"strings"

	"google.golang.org/api/chat/v1"
)

func GetMessageText(event *chat.DeprecatedEvent, toLower bool) string {
	if event.Message == nil || event.Type != "MESSAGE" {
		return ""
	}

	command := strings.TrimSpace(event.Message.ArgumentText)

	if toLower {
		command = strings.ToLower(command)
	}

	return command
}
