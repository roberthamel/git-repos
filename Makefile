build:
	@go build -o $(HOME)/bin/git-repos . && which git-repos

test:
	@go test -v ./...

coverage:
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o cover.html
	@live-server cover.html

clean:
	@rm -rf cover.html coverage.out
