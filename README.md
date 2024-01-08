# PetAutchTest

https://habr.com/ru/articles/774796/

https://github.com/GolangLessons/sso - проект

sso
├── cmd.............. Команды для запуска приложения и утилит
│   ├── migrator.... Утилита для миграций базы данных
│   └── sso......... Основная точка входа в сервис SSO
├── config........... Конфигурационные yaml-файлы
├── internal......... Внутренности проекта
│   ├── app.......... Код для запуска различных компонентов приложения
│   │   └── grpc.... Запуск gRPC-сервера
│   ├── config....... Загрузка конфигурации
│   ├── domain
│   │   └── models.. Структуры данных и модели домена
│   ├── grpc
│   │   └── auth.... gRPC-хэндлеры сервиса Auth
│   ├── lib.......... Общие вспомогательные утилиты и функции
│   ├── services..... Сервисный слой (бизнес-логика)
│   │   ├── auth
│   │   └── permissions
│   └── storage...... Слой работы с данными
│       └── sqlite.. Реализация на SQLite
├── migrations....... Миграции для базы данных
├── storage.......... Файлы хранилища, например SQLite базы данных
└── tests............ Функциональные тесты

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

Слой работы с данными — Storage
использовать SQLite
go get github.com/mattn/go-sqlite3@v1.14.17

Миграции БД
go get github.com/golang-migrate/migrate/v4@v4.16.2

будем использовать библиотеку grpc-ecosystem/go-grpc-middleware, содержащую готовые реализации некоторых полезных интерсепторов
go get github.com/grpc-ecosystem/go-grpc-middleware/v2@v2.0.0