package commands

import (
	"fmt"

	"google.golang.org/api/chat/v1"

	"github.com/iVariable/Highfive/internal/pkg/highfive"
)

type praiseWriter interface {
	SavePraise(praise highfive.Praise) error
}

type Praise struct {
	PraiseStorage praiseWriter
}

func (p Praise) Is(event *chat.DeprecatedEvent) bool {
	if event.Message == nil {
		return false
	}

	praises, err := highfive.PraiseFromMessage(*event.Message)

	return err == nil && praises != nil && len(praises) != 0
}

func (p Praise) Exec(event *chat.DeprecatedEvent) (*chat.Message, error) {
	praises, err := highfive.PraiseFromMessage(*event.Message)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse praises: %w`, err)
	}

	for _, praise := range praises {
		e := p.PraiseStorage.SavePraise(praise)
		if e != nil {
			return nil, fmt.Errorf(`failed to save praise: %w`, e)
		}
	}

	return &chat.Message{
		Text: "🎉",
	}, nil
}

func (p Praise) Help(event *chat.DeprecatedEvent) string {
	return "*Praise someone:*\n\n" +
		"- to @username: praise @username without specifying a reason\n" +
		"- to @username for {reason}: praise @username for a reason (you don't need to type brackets of course)\n" +
		"- to @username because {reason}: praise @username for a reason (you don't need to type brackets of course)\n" +
		"- to @usernameA @usernameB and @usernameC for {reason}: praise multiple people for a reason\n\n" +
		"*Example:* @Highfive to @" + event.User.DisplayName + " for using highfive help command!"
}
