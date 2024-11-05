
# Fx + Gofiber + MongoDB
## Require Swaggo
```
go install github.com/swaggo/swag/cmd/swag@1.8.7
```
## 1. RUN environment
```
    cd .development && docker compose up -d
```
## 2. Copy ENV file
```
    cp local.env.example local.env
```
## 2. Project command
```
    make docs // generate swagger doc
    make api // start project
```