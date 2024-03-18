package main

import (
	"encoding/json"
	"log"
	"sharedlambdacode/shared/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const MerchantGroup = "Merchant"
const CustomerGroup = "Customer"

type ApiAccessInfo struct {
	allowMerchant bool
	allowCustomer bool
}

var apis map[string]ApiAccessInfo = map[string]ApiAccessInfo{
	"/hello": {
		allowMerchant: true,
	},
	"/queue_processor": {
		allowCustomer: true,
	},
}

func handler(request events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
    log.Printf("Current time: %s", util.GetCurrentTimeISO())
	reqJson, err := json.Marshal(request)
	if err != nil {
		log.Printf("could not marshal json: %s", err)
	} else {
		log.Printf("request JSON: %s", string(reqJson))
	}

	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
	}, nil
}

func main() {
	lambda.Start(handler)
}
