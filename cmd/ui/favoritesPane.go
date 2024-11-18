package ui

import "github.com/charmbracelet/lipgloss"

func FavoritesPane(currentSelectedPane int) string {
	folderSelectStyle := lipgloss.NewStyle().Width(30).Border(lipgloss.NormalBorder()).Height(19)

	if currentSelectedPane == 2 {
		folderSelectStyle = folderSelectStyle.BorderForeground(HIGHLIGHT_COLOR)
	}
	return folderSelectStyle.Render("Favorites[2]") }
