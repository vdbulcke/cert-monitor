package cmd

import (
	"fmt"
	"runtime"
	rtdebug "runtime/debug"

	"github.com/spf13/cobra"
)

// GitCommit the current git commit
// will be injected during build
var GitCommit string

// Version
var Version string

// Date
var Date string

// BuiltBy
var BuiltBy string

// HumanVersion version with commit
var HumanVersion = fmt.Sprintf("%s-(%s)", Version, GitCommit)

// Args variable
var short bool

func init() {
	// bind to root command
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&short, "short", "", false, "short version info")

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cert-monitor",
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			fmt.Println(HumanVersion)

		} else {
			fmt.Println(buildVersion())

		}
	},
}

// ref: goreleaser
func buildVersion() string {
	result := fmt.Sprintf("version: %s", Version)
	if GitCommit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, GitCommit)
	}
	if Date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, Date)
	}
	if BuiltBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, BuiltBy)
	}
	result = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s", result, runtime.GOOS, runtime.GOARCH)
	if info, ok := rtdebug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
	}
	return result
}
