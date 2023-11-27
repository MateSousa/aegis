install: # Install dependencies
	@go mod tidy

start-api: # Start api 
	@go run cmd/api/main.go

start-connect: # Start connect 
	@go run cmd/connect/main.go

test: # Run all unit tests	
	@go test ./... -timeout 5s -cover -coverprofile=cover.out	

doc: # Run generate documentation
	@if [ ! -f "$(GOPATH)/bin/swag" ]; then \
		echo "Swag not found. Installing Swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@swag init --ot go -g cmd/main.go

sec: # Run security tests
	@if [ ! -f "$(GOPATH)/bin/gosec" ]; then \
		echo "Gosec not found. Installing Gosec..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
	fi
	@gosec ./...
