package gen

import (
	"fmt"

	"github.com/shiron-dev/arcanum-hue/internal/converter"
	"github.com/spf13/cobra"
)

const itermRequiredArgs = 2

var itermCmd = &cobra.Command{
	Use:   "iterm",
	Short: "Generate color theme for iTerm2",
	Long: `Generate a color theme for iTerm2.
Usage: arcanumhue gen iterm [config file path] [output file path]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen iterm called")

		if len(args) < itermRequiredArgs {
			fmt.Println(`Please specify the config file path and the output file path.
	Usage: arcanumhue gen iterm [config file path] [output file path]`)

			return
		}

		cfgPath := args[0]
		fmt.Println("config path:", cfgPath)

		outPath := args[1]
		fmt.Println("output path:", outPath)

		err := converter.GetItermTheme(cfgPath, outPath)
		if err != nil {
			fmt.Println("Failed to generate iTerm2 theme:", err)

			return
		}

		fmt.Println("Successfully generated iTerm2 theme")
	},
}

func init() {
	genCmd.AddCommand(itermCmd)
}
