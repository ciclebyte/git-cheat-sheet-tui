package main

import (
	"embed"

	"github.com/ciclebyte/git-cheat-sheet-tui/assets"
	"github.com/ciclebyte/git-cheat-sheet-tui/cmd"
)

//go:embed resources
var resources embed.FS

func main() {
	assets.Resources = resources
	cmd.Execute()
}
