GOOS=linux GOARCH=amd64 go build -o build/create-flask-app-for-linux main.go
GOOS=windows GOARCH=amd64 go build -o build/create-flask-app.exe main.go
