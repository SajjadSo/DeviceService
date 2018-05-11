package main

import (
	"DeviceService/Responses"
	"DeviceService/Types"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Handler function: Get user input parameter and returns a specific device if exists
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get request id, and check if id is valid
	ID := req.PathParameters["id"]
	if ID == "" {
		return responses.ErrNotFound(), nil
	}

	// Initialize a session in AWS_REGION that the SDK will use to load
	region := os.Getenv("AWS_REGION")
	sess := session.Must(session.NewSession(&aws.Config{Region: &region}))

	// Create DynamoDB client
	db := dynamodb.New(sess)

	// Get dbDevice from DevicesTable
	DevicesTable := aws.String(os.Getenv("DEVICES_TABLE"))
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: DevicesTable,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ID),
			},
		},
	})

	if err != nil {
		return responses.ErrInternalServer(), nil
	}

	// If length of result is zero, then return NotFound response 404
	if len(result.Item) == 0 {
		return responses.ErrNotFound(), nil
	}

	// Create empty Device object for return to response
	device := types.Device{}

	// Convert dynamodb result to Device object
	if err = dynamodbattribute.UnmarshalMap(result.Item, &device); err != nil {
		return responses.ErrInternalServer(), nil
	}

	// Return found device to response
	deviceJSON, err := json.Marshal(device)
	if err != nil {
		return responses.ErrInternalServer(), nil
	}

	return responses.SuccessDeviceFound(deviceJSON), nil
}

func main() {
	lambda.Start(Handler)
}
