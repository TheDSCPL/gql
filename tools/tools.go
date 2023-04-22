//go:build false

// To install the following tools at the version used by this repo run:
// $ make download-tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)
