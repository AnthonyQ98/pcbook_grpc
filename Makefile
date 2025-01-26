gen:
	protoc --proto_path=proto proto/*.proto --go_out=./pb --go-grpc_out=./pb --grpc-gateway_out=:pb --openapiv2_out=:openapiv2

clean:
	rm pb/*pb.go

server1:
	go run cmd/server/main.go --port 50051

server2:
	go run cmd/server/main.go --port 50052

server1-tls:
	go run cmd/server/main.go --port 50051 -tls

server2-tls:
	go run cmd/server/main.go --port 50052 -tls


server:
	go run cmd/server/main.go --port 8080

rest:
	go run cmd/server/main.go --port 8081 --type rest


client:
	go run cmd/client/main.go --address 0.0.0.0:8080

client-tls:
	go run cmd/client/main.go --address 0.0.0.0:8080 -tls

test:
	go test -cover -race ./...

cert:
	cd ca-certificate-generator; ./gen.sh; cd ..

.PHONY: gen clean server client test cert