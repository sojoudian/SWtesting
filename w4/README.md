go mod init w4
go mod tidy

go mod init github.com/sojoudian/SWtesting/w4
go mod tidy


go run main.go
Dockerfile

docker tag w4:v0.1 maziar/w4:v0.1
docker push maziar/w4:v0.1


