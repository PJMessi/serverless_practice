package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.SQSEvent) (events.APIGatewayProxyResponse, error) {
    log.Println("Queue Processor Lamabda Invoked")

    log.Println(request.Records)
    log.Println(len(request.Records))

    return events.APIGatewayProxyResponse{
        Body: "queue processor lambda processed",
        StatusCode: 200,
    }, nil
}


func main() {
	lambda.Start(handler)
}
