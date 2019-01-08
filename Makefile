GO=go

grpc:
	protoc -I grtest/ --go_out=plugins=grpc:grtest/ grtest/grtest.proto
	python3 -m grpc_tools.protoc -I grtest/ --python_out=grtest --grpc_python_out=grtest grtest/grtest.proto

go_server: grpc
	${GO} build -o go_server server/server.go
	./go_server

go_client: grpc
	${GO} build -o go_client client/client.go
	./go_client

python_server: grpc
	python3 server/server.py

python_client: grpc
	python3 client/client.py
