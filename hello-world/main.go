package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type QueueMessageType struct {
	BookName   string
	AuthorName string
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
			BookName: "Harry Potter and the Philosopher's Stone",
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
