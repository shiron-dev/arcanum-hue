package gen

import (
	"fmt"

	"github.com/shiron-dev/arcanum-hue/internal/converter"
	"github.com/spf13/cobra"
)

var vscodeCmd = &cobra.Command{
	Use:   "vscode",
	Short: "Generate color theme for Visual Studio Code",
	Long:  `Generate a color theme for Visual Studio Code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen vscode called")

		if cmd.Flag("config").Value.String() == "" {
			fmt.Println("Please specify the configuration file")

			return
		}

		cfgPath := cmd.Flag("config").Value.String()
		fmt.Println("config path:", cfgPath)

		err := converter.GetVSCodeTheme(cfgPath, "vscode-theme.json")
		if err != nil {
			fmt.Println("Failed to generate Visual Studio Code theme:", err)

			return
		}
	},
}

func init() {
	genCmd.AddCommand(vscodeCmd)

	vscodeCmd.Flags().StringP("config", "c", "", "Path to the configuration file")
}
