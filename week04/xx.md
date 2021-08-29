github.com/golang/protobuf 被 google.golang.org/protobuf  取代

go get -u google.golang.org/protobuf   安装 protoc-gen-go， 生成在 $GOPATH/bin 目录下

 安装  protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

```
obcs-MacBook-Pro:helloworld obc$ echo $GOPATH/bin
/Users/obc/Desktop/workspace/gomodpath/bin
obcs-MacBook-Pro:helloworld obc$ ls $GOPATH/bin
protoc-gen-go		protoc-gen-swagger
protoc-gen-grpc-gateway	wire

obcs-MacBook-Pro:week04 obc$ tree helloworld/
helloworld/
└── proto
    ├── helloworld.pb.go
    ├── helloworld.pb.gw.go
    └── helloworld.proto
```
```
syntax = "proto3";

// 包名
package helloworld;


import "google/api/annotations.proto";

// 定义的服务名
service Greeter {
  // 具体的远程服务方法
  rpc SayHello (HelloRequest) returns (HelloReply) {
    option (google.api.http) = {
      post: "/helloworld"
      body: "*"
    };
  }
}

// SayHello方法的入参，只有一个字符串字段
message HelloRequest {
  string name = 1;
}

// SayHello方法的返回值，只有一个字符串字段
message HelloReply {
  string message = 1;
}
```

在helloworld 目录下执行：
```
obcs-MacBook-Pro:helloworld obc$ protoc -I . --go_out=plugins=grpc:./proto ./proto/helloworld.proto  --proto_path=/Users/obc/Desktop/workspace/gomodpath/pkg/mod/github.com/go-kratos/kratos/v2@v2.0.0-rc6/third_party
```
出错：
```
protoc-gen-go: invalid Go import path "helloworldv1" for "proto/helloworld.proto"
Please specify either:
	• a "go_package" option in the .proto source file, or
	• a "M" argument on the command line.
```
See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

解决：https://blog.csdn.net/weixin_43851310/article/details/115431651
然后 在helloworld.proto 加上：option go_package = "./;helloworld";
--go_out=plugins=grpc:./proto 这个./proto 加上./ 就是生成helloworld.pb.go 的路径，
如果--go_out=plugins=grpc:./ ，那么 option go_package = "./protoc;helloworld";

生成helloworld.pb.gw.go
protoc -I . --grpc-gateway_out ./   \
  --grpc-gateway_opt logtostderr=true \
      --grpc-gateway_opt paths=source_relative \
       --proto_path=/Users/obc/Desktop/workspace/ \
       gomodpath/pkg/mod/github.com/go-kratos/kratos/v2@v2.0.0-rc6/third_party ./proto/helloworld.proto

-----------------------------------------------------------------



blog.proto
```
syntax = "proto3";

package blog.api.v1;

option go_package = "api/v1;v1";
option java_multiple_files = true;
option java_package = "blog.api.v1";

import "google/api/annotations.proto";
// the validate rules:
// https://github.com/envoyproxy/protoc-gen-validate
//import "validate/validate.proto";

`````

protoc -I . --grpc-gateway_out ./     --grpc-gateway_opt logtostderr=true     --grpc-gateway_opt paths=source_relative  --proto_path=/Users/obc/Desktop/workspace/gomodpath/pkg/mod/github.com/go-kratos/kratos/v2@v2.0.0-rc6/third_party   api/v1/blog.proto

 --grpc-gateway_out  表示调用 protoc-gen-grpc-gateway ， 即 

--proto_path  指定 blog.proto import 目录， 比如在blog.proto 里有：
import "google/api/annotations.proto";
那么 就是指定 $ {--proto_path}/google/api/annotations.proto

go_package 前部分是指定生成xx.pb.go 的路径，是相对路径， 后半部分，就生成xx.pb.go 时的包名，package v1