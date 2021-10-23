## JWT generator & parser microservices

### api service (parsing & validation)

```shell
# install dependencies for api service
cd api && go mod tidy

# run api service
go run main.go
```

### jwt_creator service (generation)

```shell
# install dependencies for jwt_creator service
cd jwt_creator && go mod tidy

# run jwt_creator service
go run main.go
```
