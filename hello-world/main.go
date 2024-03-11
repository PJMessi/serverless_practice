package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queueUrl := os.Getenv("HELLO_WORLD_SQS_QUEUE_URL")
	log.Printf("Queue URL: %s", queueUrl)

	svc := sqs.New(sess)
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(1),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": {
				DataType: aws.String("String"),
				StringValue: aws.String("Harry Potter and the Philosopher's Stone"),
			},
			"Author": {
				DataType: aws.String("String"),
				StringValue: aws.String("J.K. Rowling"),
			},
		},
		MessageBody: aws.String("Information about Books"),
		QueueUrl: aws.String(queueUrl),
	})

	if err != nil {
		log.Printf("error while publishing to queue: %v", err)
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
