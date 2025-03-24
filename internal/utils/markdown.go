package utils

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

var MarkdownInstance = Markdown{}

type Markdown struct {
	renderer *glamour.TermRenderer
}

func init() {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(180),
		glamour.WithEmoji(),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	MarkdownInstance.renderer = renderer

}

func (c *Markdown) MarkdownRender(content string) (string, error) {
	out, err := c.renderer.Render(content)
	if err != nil {
		return "", err
	}
	return out, nil
}
