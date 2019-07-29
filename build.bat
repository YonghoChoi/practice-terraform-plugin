@echo off

set GO111MODULE=on
set GOARCH=amd64
set GOOS=windows

rd /s /q .\bin

go mod tidy
go build -o .\bin\terraform-provider-example