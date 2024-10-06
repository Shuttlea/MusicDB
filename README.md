#  MUSIC DB

## Start program

`go run *.go`

## Environment

For change environment parameters find the file: `config/config.env`

#### Standart Parameters:

- HOST_ADDR=localhost:8888
- DATABASE_HOST=localhost
- DATABASE_USER=shuttlea 
- DATABASE_DB=postgres
- DATABASE_PORT=5432
- DATABASE_SSLMODE=disable

## Swagger

Swagger file is `api/swagger.yaml`

## Logger

Change Log parameters (INFO/DEBUG) you can in `main.go` file in line 22

## Second exersize

2. При добавлении сделать запрос в АПИ, описанного сваггером

Задание реализовано в файле `api_handlers.go` в строках 127-128
