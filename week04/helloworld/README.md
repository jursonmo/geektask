```
//安装 protoc-gen-go， 生成在 $GOPATH/bin 目录下
go get -u google.golang.org/protobuf   

 //安装  protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 

生成helloworld.pb.go
protoc -I . --go_out=plugins=grpc:./proto ./proto/helloworld.proto  --proto_path=`pwd`/proto_api

生成helloworld.pb.gw.go
protoc -I . --grpc-gateway_out ./     --grpc-gateway_opt logtostderr=true  \ 
--grpc-gateway_opt paths=source_relative  \
--proto_path=`pwd`/proto_api ./proto/helloworld.proto
```
