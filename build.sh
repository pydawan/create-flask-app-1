GOOS=linux GOARCH=386 go build -o build/create-flask-app-for-linux-386 main.go
GOOS=linux GOARCH=amd64 go build -o build/create-flask-app-for-linux-amd64 main.go
GOOS=windows GOARCH=386 go build -o build/create-flask-app-for-windows-386.exe main.go
GOOS=windows GOARCH=amd64 go build -o build/create-flask-app-for-windows-amd64.exe main.go
