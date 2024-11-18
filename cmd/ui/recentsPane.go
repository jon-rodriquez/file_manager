package ui

import "github.com/charmbracelet/lipgloss"

func RecentsPane(currentSelectedPane int) string {
	folderSelectStyle := lipgloss.NewStyle().Width(30).Border(lipgloss.NormalBorder()).Height(19)

	if currentSelectedPane == 3 {
		folderSelectStyle = folderSelectStyle.BorderForeground(HIGHLIGHT_COLOR)
	}
	return folderSelectStyle.Render("Recents[3]")
}
