test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

build-release:
	go build -ldflags "-s -w"

build-arm7:
	GOARCH=arm GOARM=7 go build -ldflags "-s -w"
