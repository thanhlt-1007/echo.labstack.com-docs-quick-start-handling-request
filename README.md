# echo.labstack.com-docs-quick-start-handling-request

- Handling Request

- Reference: https://echo.labstack.com/docs/quick-start#handling-request

## gvm

```sh
gvm install go1.24.0
gvm use go1.24.0
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

```sh
curl --location 'http://localhost:1323/users/create' \
--form 'name="Joe Smith"' \
--form 'email="joe.smith@gmail.com"'
```

## gofumpt

- Install

```sh
go install mvdan.cc/gofumpt@latest
```

- Run

```sh
gofumpt -l -w .
```

## golangci-lint

- Install

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
```

- Run

```sh
golangci-lint run
```
