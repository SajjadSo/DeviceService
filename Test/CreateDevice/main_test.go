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

// PutItem function for DynamoDB Mock
func (d *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	// mockDynamoDB PutItem function response
	return new(dynamodb.PutItemOutput), nil
}

// array that holds tests queue
// based on struct(request, response) for each test
var tests = []struct {
	req events.APIGatewayProxyRequest
	res int
}{
	{
		// Bad json request test
		events.APIGatewayProxyRequest{
			Body: "{123}",
		},
		400,
	},
	{
		// Missing Fields test (id, deviceModel)
		events.APIGatewayProxyRequest{
			Body: "{\"id\":\"1\", \"deviceModel\":\"deviceModel1\"}",
		},
		400,
	},
	{
		// Missing Fields test (name, note, serial)
		events.APIGatewayProxyRequest{
			Body: "{\"name\":\"name1\", \"note\":\"note1\", \"serial\":\"serial1\"}",
		},
		400,
	},
	{
		// Correct json request
		events.APIGatewayProxyRequest{
			// Headers: map[string]string{"Content-Type": "application/json"},
			Body: "{\"id\":\"1\", \"deviceModel\":\"deviceModel1\", \"name\":\"name1\", \"note\":\"note1\", \"serial\":\"serial1\"}",
		},
		201,
	},
}

// TestHandler function
func TestHandler(t *testing.T) {
	for i, testItem := range tests {
		// Send request of each test to Handler function
		response, _ := Handler(testItem.req)
		if response.StatusCode != testItem.res {
			t.Errorf("Test Fail in tests[%d]: want: %d, got: %d", i, testItem.res, response.StatusCode)
		}
	}
}
