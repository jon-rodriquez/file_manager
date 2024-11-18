package ui

import "github.com/charmbracelet/lipgloss"

func FileManagerPane(currentSelectedPane int) string {
	rightNavStyle := lipgloss.NewStyle().Width(30).Border(lipgloss.NormalBorder()).Height(2)
	if currentSelectedPane == 1 {
		rightNavStyle = rightNavStyle.BorderForeground(HIGHLIGHT_COLOR)
	}
	return rightNavStyle.Render("File Manager[1]")
}
