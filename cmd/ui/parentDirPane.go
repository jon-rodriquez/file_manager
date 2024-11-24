package ui

import (
	"file_manager/cmd/dir"

	"github.com/charmbracelet/lipgloss"
)

func ParentDirPane(currentSelectedPane int, currentDir []dir.Item, cursor int) string {
	columnStyle := lipgloss.NewStyle().Width(60).Border(lipgloss.NormalBorder()).Height(40)

	if currentSelectedPane == 5 {
		columnStyle = columnStyle.BorderForeground(HIGHLIGHT_COLOR)
	}
	return columnStyle.Render(ListView(currentDir, cursor))
}
