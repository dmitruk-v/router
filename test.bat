@echo off
go test -v -cover -coverprofile=cover.out -run=^Test github.com/dmitruk-v/router/v1
