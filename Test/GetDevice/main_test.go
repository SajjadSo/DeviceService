package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// A DynamoDB Mock
type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

// GetItem function for DynamoDB Mock
func (d *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	// mockDynamoDBClient GetItem function response
	id := input.Key["id"].S
	ItemOutput := dynamodb.GetItemOutput{}

	// If id=1, Set ItemOutput
	if *id == "1" {
		ItemOutput.SetItem(
			map[string]*dynamodb.AttributeValue{
				"id": &dynamodb.AttributeValue{S: id},
			},
		)
	}

	return &ItemOutput, nil
}

// array that holds tests queue
// based on struct(request, response) for each test
var tests = []struct {
	req events.APIGatewayProxyRequest
	res int
}{
	{
		// Field 'id' is empty
		events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "",
			},
		},
		404,
	},
	{
		// request id not found in database
		events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "123",
			},
		},
		404,
	},
	{
		// request id found in database
		events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "1",
			},
		},
		200,
	},
}

// TestHandler function
func TestHandler(t *testing.T) {
	for i, testItem := range tests {
		// Send request of each test to Handler function
		response, _ := Handler(testItem.req)
		if response.StatusCode != testItem.res {
			t.Errorf("Test Fail in tests[%d]: want: %d, got: %d \n %s", i, testItem.res, response.StatusCode, response.Body)
		}
	}
}
