CURRENT_PATH = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
HOME_PATH = $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST))))../)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

gen-model:
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/mall" -table="*" -dir="./model"

user-api:
	goctl api go -api ./service/user/api/user.api -dir ./service/user/api
