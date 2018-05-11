# Simple Restful API on AWS

Implement a simple Restful API on AWS using the following tech stack:

 * Serverless Framework ( https://serverless.com ) 
 * Go language ( https://golang.org )
 * AWS API Gateway
 * AWS Lambda
 * AWS DynamoDB


The API should accept the following JSON requests and produce the corresponding HTTP responses:

#### Request 1:
```
HTTP POST
URL: https://<api-gateway-url>/api/devices
Body (application/json):
{
  "id": "/devices/id1",
  "deviceModel": "/devicemodels/id1",
  "name": "Sensor",
  "note": "Testing a sensor.",
  "serial": "A020000102"
}
```

#### Response 1 - Success:
```
HTTP 201 Created
Body (application/json):
{
  "id": "/devices/id1",
  "deviceModel": "/devicemodels/id1",
  "name": "Sensor",
  "note": "Testing a sensor.",
  "serial": "A020000102"
}
```

#### Response 1 - Failure 1:
```
HTTP 400 Bad Request
```
If any of the payload fields are missing. Response body should have a descriptive error message for the client to be able to detect the problem.

#### Response 1 - Failure 2:
```
HTTP 500 Internal Server Error
```
If any exceptional situation occurs on the server side.

#### Request 2:
```
HTTP GET
URL: https://<api-gateway-url>/api/devices/{id}
Example: GET https://api123.amazonaws.com/api/devices/id1
```

#### Response 2 - Success:
```
HTTP 200 OK
Body (application/json):
{
  "id": "/devices/id1",
  "deviceModel": "/devicemodels/id1",
  "name": "Sensor",
  "note": "Testing a sensor.",
  "serial": "A020000102"
}
```

#### Response 2 - Failure 1:
```
HTTP 404 Not Found
```
If the request id does not exist.

#### Response 2 - Failure 2:
```
HTTP 500 Internal Server Error
```
If any exceptional situation occurs on the server side.

## Project Structure
The project has two folders with a main.go file for each Lambda function:
 1. `CreateDevice` to handle POST request
 2. `GetDevice` to handle GET request

 Test files and their results are in the `Test` Folder.

`serverless.yml` file is responsible to manage custom values, project packaging options, resources and permissions. The serverless will use that to deploy the project to AWS.

because of the project is developed on Windows, the commands for building project saved in `Makefile` file.

## Deployment

### Prerequisites
Before deploy with Serverless, we must install the following tools:
(The environment of my development is on Windows 10)

 * aws cli: i use aws msi installer -> https://s3.amazonaws.com/aws-cli/AWSCLI64.msi
 * go: i use https://dl.google.com/go/go1.10.1.windows-amd64.msi from https://golang.org/dl/ webpage
 * go dep
 * build-lambda-zip: go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip and go install
 * nodejs: i use https://nodejs.org/dist/v8.11.1/node-v8.11.1-x64.msi from https://nodejs.org/en/download/ webpage
 * serverless: in CMD run `npm install -g serverless` command

### Deploy
 * In CMD go to folder `GOPATH\src`
 * Enter command `sls create --template aws-go-dep --path MyServiceName`
 * Create folder for each function, with main.go file in that
 * Create folder `bin`
 * To build functions use the following commands in each folder of lambda functions :
   - %USERPROFILE%\Go\bin\dep ensure
   - set GOOS=linux
   - go build -o bin/CreateDevice CreateDevice/main.go
   - go build -o bin/GetDevice GetDevice/main.go
   - %USERPROFILE%\Go\bin\build-lambda-zip.exe -o bin/CreateDevice.zip bin/CreateDevice
   - %USERPROFILE%\Go\bin\build-lambda-zip.exe -o bin/GetDevice.zip bin/GetDevice
 * I use `build-lambda-zip` to package compiled lambda functions, and then set `artifact` in `serverless.yml` to force serverless use these packages for uploading functions
 * The executable files will be created in `bin` folder, after you run.
 * In order to deploy, run `sls deploy`

After deploy successfully, two URLs will be generated. These are api URLs called `API-GATEWAY-URL`

## Testing
For testing project, i use `Postman` windows software which can be download on https://www.getpostman.com/apps

### Test `CreateDevice` function
in `Postman` enter URL `https://<api-gateway-url>/api/devices` and choose `POST` request, and in `Body` tab choose `raw` radiobutton and enter the following json request body:
```
{
  "id": "/devices/id1",
  "deviceModel": "/devicemodels/id1",
  "name": "Sensor",
  "note": "Testing a sensor.",
  "serial": "A020000102"
}
```
then click on `Send` button.
The request will send to URL and show response in bottom section.

### Test `GetDevice` function
in `Postman` enter URL `https://<api-gateway-url>/api/devices/{id}` and choose `GET` request (replace {id}  with `id1` or something else), then click on `Send` button.
The request will send to URL and show response in bottom section.

### Unit test coverage
In `Test` folder, go to each folder of functions and open Command Prompt and use the following commands to generate cover output and its HTML view:
```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Here is the results of my test: 
#### CreateDevice:
```
PASS
coverage: 88.9% of statements
ok      DeviceService/Test/CreateDevice 0.108s
```

#### GetDevice:
```
PASS
coverage: 80.0% of statements
ok      DeviceService/Test/GetDevice    0.115s
```
