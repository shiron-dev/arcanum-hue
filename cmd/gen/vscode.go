package gen

import (
	"fmt"

	"github.com/shiron-dev/arcanum-hue/internal/converter"
	"github.com/spf13/cobra"
)

var vscodeCmd = &cobra.Command{
	Use:   "vscode",
	Short: "Generate color theme for Visual Studio Code",
	Long: `Generate a color theme for Visual Studio Code.
Usage: arcanumhue gen vscode [config file path] [output file path]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen vscode called")

		if len(args) < 2 {
			fmt.Println(`Please specify the config file path and the output file path.
	Usage: arcanumhue gen vscode [config file path] [output file path]`)

			return
		}

		cfgPath := args[0]
		fmt.Println("config path:", cfgPath)

		outPath := args[1]
		fmt.Println("output path:", outPath)

		err := converter.GetVSCodeTheme(cfgPath, outPath)
		if err != nil {
			fmt.Println("Failed to generate Visual Studio Code theme:", err)

			return
		}
	},
}

func init() {
	genCmd.AddCommand(vscodeCmd)
}
