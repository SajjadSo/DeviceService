package responses

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// ErrBadRequest function: return bad request payload error
func ErrBadRequest(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       "Bad Request\n" + err.Error(),
	}
}

// ErrInternalServer function: return internal server error
func ErrInternalServer() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       "Internal Server Error",
	}
}

// ErrNotFound function: return not found error
func ErrNotFound() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       "Not Found",
	}
}

// SuccessCreated function: return response for create a new device successfully
func SuccessCreated(deviceJSON []byte) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(deviceJSON),
	}
}

// SuccessDeviceFound function: return response for found a device successfully
func SuccessDeviceFound(deviceJSON []byte) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(deviceJSON),
	}
}
