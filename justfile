default: ci

update:
    go get -u

tidy:
    go mod tidy

test:
    go test ./...

toolchain:
    go install honnef.co/go/tools/cmd/staticcheck@latest
    go install github.com/kisielk/errcheck@latest

staticcheck: toolchain
    staticcheck -f stylish ./...
    errcheck ./...

init: tidy update

vendor:
    go mod vendor

vet:
    go vet ./...

ci: init vendor vet staticcheck test
