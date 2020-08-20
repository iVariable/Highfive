package cmd

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Response(code int, text string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: text,
	}
}

func CodeResponse(code int) events.APIGatewayProxyResponse {
	return Response(code, http.StatusText(code))
}

func JSONResponse(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",

			"Content-Type": "application/json",
		},
		Body: body,
	}
}

func MarshalJSONResponse(body interface{}) (events.APIGatewayProxyResponse, error) {
	response, err := json.Marshal(body)
	if err != nil {
		return CodeResponse(http.StatusInternalServerError), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",

			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func TextResponse(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",

			"Content-Type": "text/plain",
		},
		Body: body,
	}
}

func Render(w http.ResponseWriter, response events.APIGatewayProxyResponse) {
	for name, value := range response.Headers {
		w.Header().Set(name, value)
	}

	w.WriteHeader(response.StatusCode)

	//nolint:errcheck
	w.Write([]byte(response.Body))
}

func BinaryResponse(res []byte, responseType string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Credentials": "true",

			"Content-Type": responseType,
		},
		IsBase64Encoded: true,
		Body:            base64.StdEncoding.EncodeToString(res),
	}
}
