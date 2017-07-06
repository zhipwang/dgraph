#!/bin/bash

protos=$GOPATH/src/github.com/dgraph-io/dgraph/protos
pushd $protos > /dev/null
protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--gofast_out=plugins=grpc:. \
	*.proto
protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	*.proto
