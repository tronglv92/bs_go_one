include .env
OS := $(shell uname)
run:
ifeq ($(OS),Darwin)
	@echo "Setting environment variables from .env on macOS"
	@set -o allexport && source .env && set +o allexport && go run -mod=vendor cmd/main/main.go
else
ifeq ($(OS),Linux)
	@echo "Setting environment variables from .env on Linux"
	@export $(shell cat .env | xargs) && go run -mod=vendor cmd/main/main.go
else
	@echo "Unsupported operating system: $(OS)"
endif
endif