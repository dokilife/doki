#!/bin/bash

base_dir=$(pwd)

gen_proto_go() {
    echo "Generating Go protobuf files for $1"
    cd "$base_dir"/src/"$1" || return
    protoc -I ../../pb ./../../pb/demo.proto --go_out=./ --go-grpc_out=./
    cd "$base_dir" || return
}

gen_proto_go userservice