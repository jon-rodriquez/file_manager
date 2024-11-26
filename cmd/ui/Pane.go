package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Pane(title string, content string, width int, height int, isSelected bool) string {
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Height(height).
		BorderTop(false)

	customBorder := lipgloss.Border{
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		BottomLeft:  "└",
		BottomRight: "┘",
	}
	borderStyle = borderStyle.BorderStyle(customBorder)
	if isSelected {
		borderStyle = borderStyle.BorderForeground(HIGHLIGHT_COLOR)
	}
	topstyle := lipgloss.NewStyle()
	if isSelected {
		topstyle = topstyle.Foreground(HIGHLIGHT_COLOR)
	}
	return lipgloss.JoinVertical(lipgloss.Left, topstyle.Render(topBorder(title, width)), borderStyle.Width(width).Render(content))
}

func topBorder(title string, width int) string {
	titleWidth := lipgloss.Width(title)
	return "┌ " + title + " " + strings.Repeat("─", width-titleWidth-lipgloss.Width("┌")-1) + "┐"
}
