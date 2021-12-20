package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// GitCommit the current git commit
// will be injected during build
var GitCommit string

// Version
var Version string

// HumanVersion version with commit
var HumanVersion = fmt.Sprintf("%s-(%s)", Version, GitCommit)

func init() {
	// bind to root command
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cert-monitor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(HumanVersion)
		os.Exit(0)
	},
}
