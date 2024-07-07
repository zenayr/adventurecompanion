build:
	echo "Building Go Executable..."
	go build -o bin/adventurecompanion cmd/main/main.go
	echo "Build Complete"

run:
	./bin/adventurecompanion