package gen

import (
	"fmt"

	"github.com/shiron-dev/arcanum-hue/cmd"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate color theme for target tool",
	Long: `Generate a color theme for the target tool.
Usage: arcanumhue gen [tool]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(genCmd)
}
