package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sharedlambdacode/internal/business/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type QueueMessageType struct {
	BookName   string
	AuthorName string
}

var authService auth.AuthService = auth.NewAuthService()

func generateToken() {
	email := "pjmessi25@icloud.com"
	password := "Password123!"

	awsSession, err := session.NewSession()
	if err != nil {
		log.Printf("error creating aws session: %s", err)
		return
	}

	authFlow := cognitoidentityprovider.AuthFlowTypeUserPasswordAuth
	clientId := "bkif73f825s46ioa0ombbfu9f"
	cognitoProvider := cognitoidentityprovider.New(awsSession)
	authRes, err := cognitoProvider.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: &authFlow,
		ClientId: &clientId,
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		}),
	})

	if err != nil {
		log.Printf("error initiating auth from cognito: %s", err)
		log.Println(err)
		log.Println("===")
		return
	}

	accessToken := authRes.AuthenticationResult.AccessToken
	log.Printf("here is the access token: %s", *accessToken)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queueUrl := os.Getenv("HELLO_WORLD_SQS_QUEUE_URL")
	log.Printf("Queue URL: %s", queueUrl)

	var queueData []QueueMessageType = []QueueMessageType{
		{
			BookName:   "Harry Potter and the Philosopher's Stone",
			AuthorName: "J.K. Rowling",
		},
	}
	jsonData, err := json.Marshal(queueData)
	if err != nil {
		log.Printf("error marshaling queue data: %s", err)
	} else {
		svc := sqs.New(sess)
		_, err = svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(1),
			MessageBody:  aws.String(string(jsonData)),
			QueueUrl:     aws.String(queueUrl),
		})

		if err != nil {
			log.Printf("error while publishing to queue: %v", err)
		} else {
			log.Println("published to queue successfully")
		}
	}

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!\n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
