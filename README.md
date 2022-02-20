генератор сервера:

proto:
	protolint -fix  protos
	protoc \
	--proto_path=protos \
	--go_out=. \
	--go-grpc_out=. \
	protos/*.proto
  
  генерация клиента:
  
  clients:
	protolint -fix  protos
	protoc \
    --proto_path=protos \
	--js_out=import_style=commonjs,binary:./client     \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client \
	protos/*.proto
  
  для клиента:
  
  npm install
  
  npx webpack client.js
  
  для запуска клиента запускаем http сервер из директории клиента:
  
  python3 -m http.server 8081
  
  запуск envoy из docker образа:
  
  docker build -t my-envoy:1.0 .
  
  docker run -d -p 8080:8080 -p 9901:9901  my-envoy:1.0
  
  go run server.go
  
  доступ по адресу:
  http://localhost:8081 
