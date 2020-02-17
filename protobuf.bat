@echo off

protoc --proto_path=. --micro_out=. --go_out=. protobuf/*.proto

echo make proto files is ok!
