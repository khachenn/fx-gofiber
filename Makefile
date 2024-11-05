BINARY_FILE = sturdy-winner-api

clean:
	@rm -rf $(BINARY_FILE) coverage.out

build:
	@go build
	@ls -alh $(BINARY_FILE)

cover:
	@go test ./... -cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out

@docs:
	@swag init --instanceName $(NAME) --ot go -d ./internal/$(NAME),./lib -g ./route/swagger.route.go -o ./internal/$(NAME)/doc
	@swag fmt -g ./internal/$(NAME)/route/swagger.route.go

docs:
	@make @docs NAME=API

api: docs
	@go run main.go api || true
