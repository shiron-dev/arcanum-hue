package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var vscodeCmd = &cobra.Command{
	Use:   "vscode",
	Short: "Generate color theme for Visual Studio Code",
	Long:  `Generate a color theme for Visual Studio Code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen vscode called")
	},
}

func init() {
	genCmd.AddCommand(vscodeCmd)
}
