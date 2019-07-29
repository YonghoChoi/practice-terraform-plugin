@echo off

set GO111MODULE=on
set GOARCH=amd64
set GOOS=windows

go mod tidy
go build -o .\terraform-provider-example.exe