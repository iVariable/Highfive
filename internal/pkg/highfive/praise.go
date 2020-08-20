package highfive

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"google.golang.org/api/chat/v1"
)

type Praise struct {
	ToID string `json:"to_id"`
	To   string `json:"to"`
	//ToAvatarURL string

	FromID string `json:"from_id"`
	From   string `json:"from"`
	//FromAvatarURL string
	Reason string `json:"reason"`

	DomainID  string `json:"domain_id"`
	SpaceID   string `json:"space_id"`
	SpaceName string `json:"space_name"`
	ThreadID  string `json:"thread_id"`

	CreatedAt time.Time `json:"created_at"`
}

func PraiseFromMessage(msg chat.Message) ([]Praise, error) {
	command := msg.Text

	commandWithoutMentions := replaceAnnotations(command, msg.Annotations)

	withoutExplanation := regexp.MustCompile(`(?si)^\s*to ((@[a-z.]*)+((\s*,\s*|\s*and\s*|\s*)(@[a-z.]*)+)*)$`)
	withExplanation := regexp.MustCompile(`(?si)^\s*to ((@[a-z.]*)+((\s*,\s*|\s*and\s*|\s*)(@[a-z.]*)+)*)\s+(for|because)\s+(.+)$`)

	withTrimmedName := strings.TrimPrefix(commandWithoutMentions, "@Highfive")

	if !(withoutExplanation.MatchString(withTrimmedName) || withExplanation.MatchString(withTrimmedName)) {
		return nil, nil
	}

	explanation := ""
	lastTargetMentionIdx := len(command)

	createdAt, err := time.Parse(time.RFC3339Nano, msg.CreateTime)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse create time: %w`, err)
	}

	if withExplanation.MatchString(withTrimmedName) {
		matches := withExplanation.FindStringSubmatchIndex(withTrimmedName)
		lastTargetMentionIdx = matches[len(matches)-2]
		explanation = withTrimmedName[lastTargetMentionIdx:]

		if len(withTrimmedName) != len(commandWithoutMentions) {
			lastTargetMentionIdx += len("@Highfive")
		}
	}

	praises := []Praise{}

	for _, annotation := range msg.Annotations {
		// skip bot name call
		if annotation.StartIndex == 0 {
			continue
		}

		// Skip non mention annotations
		if annotation.Type != "USER_MENTION" ||
			annotation.UserMention == nil ||
			annotation.UserMention.User == nil ||
			annotation.UserMention.Type != "MENTION" {
			continue
		}

		//Skip mentions inside of the reason
		if int(annotation.StartIndex) > lastTargetMentionIdx {
			continue
		}

		praises = append(praises, Praise{
			ToID:      annotation.UserMention.User.Name,
			To:        annotation.UserMention.User.DisplayName,
			FromID:    msg.Sender.Name,
			From:      msg.Sender.DisplayName,
			Reason:    explanation,
			DomainID:  msg.Sender.DomainId,
			SpaceID:   msg.Space.Name,
			SpaceName: msg.Space.DisplayName,
			ThreadID:  msg.Thread.Name,
			CreatedAt: createdAt,
		})
	}

	return praises, nil
}

func replaceAnnotations(txt string, annotations []*chat.Annotation) string {
	txtOut := []rune(txt)
	replacementRune := []rune("A")[0]

	for _, an := range annotations {
		// skip bot name call
		if an.StartIndex == 0 {
			continue
		}

		for i := an.StartIndex + 1; i < an.Length+an.StartIndex; i++ {
			txtOut[i] = replacementRune
		}
	}

	return string(txtOut)
}
