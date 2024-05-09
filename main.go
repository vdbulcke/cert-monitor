package main

import (
	cmd "github.com/vdbulcke/cert-monitor/src/cmd"
)

func main() {
	// cobra commands
	// Version flags in cmd/version.go
	cmd.Execute()
}
