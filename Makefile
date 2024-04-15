test:
	set -euo pipefail
	go test -coverprofile cover.out -json -v ./... 2>&1 | tee /tmp/gotest.log | gotestfmt

lint:
	golangci-lint --color=always run ./... --fix -v
