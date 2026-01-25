Модули
go mod init <MODULE_PATH>     # например: github.com/glinozem/golang-scl-less-2/modules
go mod tidy                   # синхронизирует go.mod/go.sum (убирает лишнее, добавляет нужное)
go mod download               # скачивает зависимости в кэш
go mod vendor                 # кладёт зависимости в папку vendor

Зависимости
go get github.com/google/uuid@v1.6.0
go get -u github.com/google/uuid

Сборка
go build -o my-app ./cmd/app

Запуск
go run .
go run ./cmd/app

Тесты
go test ./...
go test -v ./...
go test -cover ./...
