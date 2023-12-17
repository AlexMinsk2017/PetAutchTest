# PetAutchTest

https://habr.com/ru/articles/774796/

библиотеки для auth
https://grpc.io/docs/languages/go/quickstart/

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ export PATH="$PATH:$(go env GOPATH)/bin"

Установка protoc
https://github.com/protocolbuffers/protobuf/releases


необходимо установить утилиту Task.
https://taskfile.dev/installation/

Для парсинга конфига-файла я буду использовать библиотеку cleanenv,
go get github.com/ilyakaznacheev/cleanenv@v1.5.0

go get github.com/GolangLessons/protos
go get golang.org/x/crypto@v0.13.0

Для работы с JWT мы будем использовать следующую библиотеку
go get github.com/golang-jwt/jwt/v5@v5.0.0