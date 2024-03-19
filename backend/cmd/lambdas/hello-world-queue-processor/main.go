package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type QueueMessageType struct {
	BookName   string
	AuthorName string
}

func handler(request events.SQSEvent) (events.APIGatewayProxyResponse, error) {
    recordsLen := len(request.Records)
    if recordsLen == 0 {
	log.Println("empty records in the request")
    } else {
	log.Printf("found %d records in the request", recordsLen)
	for i, record := range request.Records {
	    log.Printf("Processing index %d request", i)
	    messageBody := record.Body

	    var messageData []QueueMessageType
	    err := json.Unmarshal([]byte(messageBody), &messageData)
	    if err != nil {
		log.Printf("error unmarshalling the request data: %s", err)
	    }
    
	    log.Printf("found %d data inside the message body", len(messageData))
	    for _, data := range messageData {
		log.Printf("Book Name: %s", data.BookName)
		log.Printf("Author Name: %s", data.AuthorName)
	    }
	}
    }

    return events.APIGatewayProxyResponse{
        Body: "queue processor lambda processed",
        StatusCode: 200,
    }, nil
}


func main() {
	lambda.Start(handler)
}
