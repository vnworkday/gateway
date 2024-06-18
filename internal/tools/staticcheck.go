//go:build tools

package tools

// This file imports packages that are used when running static check tools. Typically, when running `make check`
import (
	_ "golang.org/x/telemetry"
	_ "golang.org/x/vuln"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
