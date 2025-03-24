package cmd

import (
	"fmt"
	"os"

	"github.com/ciclebyte/git-cheat-sheet-tui/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitsheet",
	Short: "A TUI for Git cheat sheet",
	Long:  `A terminal user interface for Git cheat sheet`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取搜索关键字
		searchKeyword, _ := cmd.Flags().GetString("search")
		// 是否显示日志
		show, _ := cmd.Flags().GetBool("showlog")
		main := app.SheetView{}
		main.Run(searchKeyword, show)
	},
}

func Execute() {
	rootCmd.Flags().StringP("search", "s", "", "Search for a command")
	rootCmd.Flags().BoolP("showlog", "l", false, "Show log panel")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nThank you for using Git Cheat Sheet TUI!, have a nice day!")
	fmt.Println("Made with ❤️ by CicleByte")
}
