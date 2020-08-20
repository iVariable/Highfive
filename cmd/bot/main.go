package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/chat/v1"

	"github.com/iVariable/Highfive/internal/cmd"
	"github.com/iVariable/Highfive/internal/cmd/bot"
	"github.com/iVariable/Highfive/internal/cmd/bot/commands"
	"github.com/iVariable/Highfive/internal/pkg/highfive"
)

var version string

func main() {
	logger := logrus.New()

	//logger.Level = logrus.InfoLevel
	logger.Level = logrus.DebugLevel

	log := logger.WithField("version", version)

	awsCfg := aws.NewConfig()
	ddb := dynamodb.New(
		session.Must(session.NewSession()),
		awsCfg,
	)

	storage := &highfive.Storage{
		Table: getEnv("STORAGE_TABLE_NAME", "Highfive"),
		DDB:   ddb,
	}

	app := &bot.Application{
		Log: log,
		Commands: []bot.Command{
			&commands.WelcomeToTheRoom{},
			&commands.Praise{
				PraiseStorage: storage,
			},
		},
		UnknownCommand: &commands.SorryIDontUnderstand{},
	}

	awslambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Debugf("%+v", req)

		event := &chat.DeprecatedEvent{}

		err := json.Unmarshal([]byte(req.Body), event)
		if err != nil {
			log.WithError(err).Warn(`failed to unmarshal bot payload`)

			return cmd.CodeResponse(http.StatusBadRequest), nil
		}

		msg, err := app.ProcessEvent(event)
		if err != nil {
			log.WithError(err).Error(`failed to marshal bot response`)

			return cmd.CodeResponse(http.StatusInternalServerError), nil
		}

		return cmd.MarshalJSONResponse(msg)
	})
}

func getEnv(name string, defaultValue string) string {
	if value, found := os.LookupEnv(name); found {
		return value
	}

	return defaultValue
}
