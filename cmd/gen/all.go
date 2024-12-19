package gen

import (
	"fmt"

	"github.com/shiron-dev/arcanum-hue/internal/converter"
	"github.com/spf13/cobra"
)

const allRequiredArgs = 2

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate color themes",
	Long: `Generate a color themes.
Usage: arcanumhue gen all [config file path] [output file path]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen all called")

		if len(args) < allRequiredArgs {
			fmt.Println(`Please specify the config file path and the output file path.
	Usage: arcanumhue gen all [config file path] [output file path]`)

			return
		}

		cfgPath := args[0]
		fmt.Println("config path:", cfgPath)

		outPath := args[1]
		fmt.Println("output path:", outPath)

		err := converter.GetAllTheme(cfgPath, outPath)
		if err != nil {
			fmt.Println("Error generating theme:", err)
		}

		fmt.Println("Theme generated successfully!")
	},
}

func init() {
	genCmd.AddCommand(allCmd)
}
