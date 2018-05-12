package main

import (
	"DeviceService/Responses"
	"DeviceService/Types"
	"encoding/json"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// sess Global variable, holds initialized session value
var sess *session.Session

// DevicesTable Global variable, holds devices table name value
var DevicesTable *string

// CheckMissingFields function: check if any fields of request payload is missing
func CheckMissingFields(PayloadJSON map[string]interface{}) error {
	Flag := false
	ErrorText := "Missing fields:"

	if PayloadJSON["id"] == nil {
		Flag = true
		ErrorText += " 'id' "
	}
	if PayloadJSON["deviceModel"] == nil {
		Flag = true
		ErrorText += " 'deviceModel' "
	}
	if PayloadJSON["name"] == nil {
		Flag = true
		ErrorText += " 'name' "
	}
	if PayloadJSON["note"] == nil {
		Flag = true
		ErrorText += " 'note' "
	}
	if PayloadJSON["serial"] == nil {
		Flag = true
		ErrorText += " 'serial' "
	}

	// If any fields are missing, return error
	if Flag == true {
		return errors.New(ErrorText)
	}

	return nil
}

// BindPayloadToDevice function: Bind request json payload to device object
func BindPayloadToDevice(PayloadJSON map[string]interface{}) types.Device {
	return types.Device{
		ID:          PayloadJSON["id"].(string),
		DeviceModel: PayloadJSON["deviceModel"].(string),
		Name:        PayloadJSON["name"].(string),
		Note:        PayloadJSON["note"].(string),
		Serial:      PayloadJSON["serial"].(string),
	}
}

// dbInsertDevice function: Insert request payload to database
func dbInsertDevice(device types.Device) error {
	// Create DynamoDB client
	db := dynamodb.New(sess)

	// Convert Device object to dynamodb AttributeValues
	dbDevice, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return err
	}

	// Insert dbDevice into DevicesTable
	_, err = db.PutItem(&dynamodb.PutItemInput{
		Item:      dbDevice,
		TableName: DevicesTable,
	})

	if err != nil {
		return err
	}

	return nil
}

// Handler function: persists the posted json Device and returns it
// back to the client as an acknowledgement
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Unmarshal request payload, and check if payload is valid
	var PayloadJSON map[string]interface{}
	if err := json.Unmarshal([]byte(req.Body), &PayloadJSON); err != nil {
		return responses.ErrBadRequest(errors.New("")), nil
	}

	// Check request payload for missing fields
	if err := CheckMissingFields(PayloadJSON); err != nil {
		return responses.ErrBadRequest(err), nil
	}

	// Bind request payload to device struct
	device := BindPayloadToDevice(PayloadJSON)

	// Insert device object to DB
	if err := dbInsertDevice(device); err != nil {
		return responses.ErrInternalServer(), nil
	}

	// Return inserted device to response
	deviceJSON, err := json.Marshal(device)
	if err != nil {
		return responses.ErrInternalServer(), nil
	}

	return responses.SuccessCreated(deviceJSON), nil
}

func main() {
	lambda.Start(Handler)
}

func init() {
	// Initialize a session in AWS_REGION that the SDK will use to load
	region := os.Getenv("AWS_REGION")
	sess = session.Must(session.NewSession(&aws.Config{Region: &region}))

	// Get Devices table name
	DevicesTable = aws.String(os.Getenv("DEVICES_TABLE"))
}
