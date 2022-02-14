//go:build tools

package axiomsegementwebhook

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "gotest.tools/gotestsum"
)
