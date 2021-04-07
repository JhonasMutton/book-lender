build:
	go mod download
	go get -u github.com/google/wire/cmd/wire@v0.5.0
	go mod vendor
	wire
	CGO_ENABLED=0 GOOS=linux go build -o bin/application

test:
	export TESTCONTAINERS_RYUK_DISABLED=true
	go mod vendor
	go get github.com/google/wire/cmd/wire@v0.5.0
	wire
	go get -v
	go test -mod=vendor -coverprofile=coverage.out ./... -json > report.out

image:
	docker build -t book-lender:latest .
