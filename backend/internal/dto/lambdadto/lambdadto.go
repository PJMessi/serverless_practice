package lambdadto

import (
	"encoding/json"
	"sharedlambdacode/internal/dto/excepdto"

	"github.com/aws/aws-lambda-go/events"
)

func ErrorToApiGatewayProxyResponse(err error) (events.APIGatewayProxyResponse, error) {
	httpExcep := excepdto.ErrorToHttpExcep(err)
	httpExcepBytes, err := json.Marshal(httpExcep)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: httpExcep.Status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(httpExcepBytes),
	}, nil
}
