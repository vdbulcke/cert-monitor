package main

import (
	cmd "github.com/vdbulcke/cert-monitor/cmd"
)

func main() {
	// cobra commands
	// Version flags in cmd/version.go
	cmd.Execute()
}
