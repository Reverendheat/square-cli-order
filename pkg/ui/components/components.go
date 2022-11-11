package components

import (
	"github.com/charmbracelet/lipgloss"
)

func Footer() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(100)
	return style.Render("Q, Quit | BackSpace, Previous")
}
