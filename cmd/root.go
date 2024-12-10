package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "arcanum-hue",
	Short: "ArcanumHue is a cli tool for integrated management of color themes",
	Long: `ArcanumHue is a cli tool for integrated management of color themes.
With ArcanumHue, you can manage color themes for various tools all at once.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoinits
func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
