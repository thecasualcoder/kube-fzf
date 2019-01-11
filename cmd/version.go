package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildVersion string

// SetVersion set the major and minor version
func SetVersion(version string) {
	buildVersion = version
}

// VersionCmd prints the build version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildVersion)
	},
}
