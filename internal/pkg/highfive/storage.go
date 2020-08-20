package highfive

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ddbWriter interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	QueryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error
}

type Storage struct {
	Table string
	DDB   ddbWriter
}

func (s *Storage) SavePraise(praise Praise) error {
	marshalledPraise, err := json.Marshal(praise)
	if err != nil {
		return fmt.Errorf(`failed to marshal praise: %w`, err)
	}

	_, err = s.DDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.Table),
		Item: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String("Highfive#" + praise.DomainID + "#" + praise.ToID),
			},
			"SK": {
				S: aws.String(strconv.FormatInt(praise.CreatedAt.UnixNano(), 10)),
			},
			"Org": {
				S: aws.String(praise.DomainID),
			},
			"Praise": {
				S: aws.String(string(marshalledPraise)),
			},
			"To": {
				S: aws.String(praise.ToID),
			},
			"From": {
				S: aws.String(praise.FromID),
			},
			"Space": {
				S: aws.String(praise.SpaceID),
			},
		},
	})

	return err
}
