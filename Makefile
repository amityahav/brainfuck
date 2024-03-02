# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build

# Name of the binary
BINARY_NAME = bf

# Main target, build the binary
build-darwin:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINARY_NAME) -v

build-linux:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINARY_NAME) -v
