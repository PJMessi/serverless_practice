package main

import (
	"encoding/json"
	"log"
	"sharedlambdacode/internal/business/auth"
	"sharedlambdacode/internal/dto/lambdadto"
	"sharedlambdacode/internal/helper/lambdahelper"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var authService auth.AuthService

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if authService == nil {
		log.Println("initialized auth service")
		authService = auth.NewAuthService()
	}
	authService = auth.NewAuthService()

	var reqBody struct {
		Phone     string `json:"phone" validate:"required"`
		EncodedPw string `json:"encodedPw" validate:"required"`
	}
	err := lambdahelper.MapRequestBody(request.Body, &reqBody)
	if err != nil {
		return lambdadto.ErrorToApiGatewayProxyResponse(err)
	}

	signInRes, err := authService.PhoneSignIn(reqBody.Phone, reqBody.EncodedPw)
	if err != nil {
		return lambdadto.ErrorToApiGatewayProxyResponse(err)
	}

	jsonRes, err := json.Marshal(signInRes)
	if err != nil {
		return lambdadto.ErrorToApiGatewayProxyResponse(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonRes),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
