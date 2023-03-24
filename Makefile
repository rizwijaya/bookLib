.PHONY : format install build

#Website/Restfull API
run:
	@echo "Running server..."
	go run ./app/main.go

init:
	@echo "Initializing dependencies..."
	go mod init
	go mod tidy
	
install:
	@echo "Downloading dependencies..."
	go mod download

build:
	@echo "building binary..."
	go build ./app/main.go

start:
	@echo "Starting server..."
	./app/main

clean:
	@echo "Cleaning..."
	rm -rf ./app/main.exe
# live reload using nodemon: npm -g i nodemon
run-nodemon:
	@echo "Running server with nodemon..."
	nodemon --exec go run ./app/main.go