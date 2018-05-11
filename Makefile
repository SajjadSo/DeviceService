build:
	%USERPROFILE%\Go\bin\dep ensure
	set GOOS=linux
	go build -o bin/CreateDevice CreateDevice/main.go
	go build -o bin/GetDevice GetDevice/main.go
	%USERPROFILE%\Go\bin\build-lambda-zip.exe -o bin/CreateDevice.zip bin/CreateDevice
	%USERPROFILE%\Go\bin\build-lambda-zip.exe -o bin/GetDevice.zip bin/GetDevice
