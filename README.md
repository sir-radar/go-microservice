## Generate code from proto file

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto`

## Build and tag docker file for upload to dockerhub

`docker build -f mail-service.dockerfile -t sirradar/mail-service:1.0.0 .`

## Push tp dockerhub

`docker push sirradar/mail-service:1.0.0`
