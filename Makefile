build:
	go build -buildmode=plugin -o=${LINTER_FOLDER}/expectfirst.so plugin/expectfirst.go

test:
	cd testdata/src/a; go mod vendor
	go test -v ./...