package highfive

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/chat/v1"
)

func TestPraiseFromMessage(t *testing.T) {
	msg := chat.Message{
		CreateTime: "2020-08-20T14:54:24.030303Z",
		Annotations: []*chat.Annotation{
			{
				Type:       "USER_MENTION",
				StartIndex: 0,
				Length:     9,
				UserMention: &chat.UserMentionMetadata{
					User: &chat.User{
						Name:        "users/111008965630742628920",
						DisplayName: "Highfive",
						Type:        "BOT",
					},
					Type: "MENTION",
				},
			},
			{
				Type:       "USER_MENTION",
				StartIndex: 13,
				Length:     20,
				UserMention: &chat.UserMentionMetadata{
					User: &chat.User{
						Name:        "users/102622061698951279126",
						DisplayName: "Yaroslav Tikhomirov",
						Type:        "HUMAN",
						DomainId:    "35gpq1j",
					},
					Type: "MENTION",
				},
			},
		},
		Sender: &chat.User{
			Name:        "users/101860392295016551101",
			DisplayName: "Vladimir Savenkov",
			Type:        "HUMAN",
			DomainId:    "35gpq1j",
		},
		Space: &chat.Space{
			Name:        "spaces/AAAAZ882Wts",
			Type:        "ROOM",
			DisplayName: "HighFive bot test room",
		},
		Text: "@Highfive to @Yaroslav Tikhomirov for not giving a fuck!",
		Thread: &chat.Thread{
			Name: "spaces/AAAAZ882Wts/threads/lxbHerdgbfk",
		},
	}

	createdAt, err := time.Parse(time.RFC3339Nano, "2020-08-20T14:54:24.030303Z")
	assert.NoError(t, err)

	expected := []Praise{
		{
			ToID:      "users/102622061698951279126",
			To:        "Yaroslav Tikhomirov",
			FromID:    "users/101860392295016551101",
			From:      "Vladimir Savenkov",
			Reason:    "not giving a fuck!",
			DomainID:  "35gpq1j",
			SpaceID:   "spaces/AAAAZ882Wts",
			SpaceName: "HighFive bot test room",
			ThreadID:  "spaces/AAAAZ882Wts/threads/lxbHerdgbfk",
			CreatedAt: createdAt,
		},
	}

	praises, err := PraiseFromMessage(msg)

	assert.NoError(t, err)
	assert.Equal(t, expected, praises)

}
