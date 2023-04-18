build:
	go build -buildmode=plugin plugin/expectfirst.go

test:
	cd testdata/src/a; go mod vendor
	go test -v ./...