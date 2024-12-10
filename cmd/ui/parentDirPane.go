package ui

import (
	"file_manager/cmd/components"
	"file_manager/cmd/types"

	"github.com/charmbracelet/lipgloss"
)

func ParentDirPane(currentSelectedPane int, currentDir []types.Item, cursor int) string {
	columnStyle := lipgloss.NewStyle().Width(60).Border(lipgloss.NormalBorder()).Height(40)

	if currentSelectedPane == 5 {
		columnStyle = columnStyle.BorderForeground(HIGHLIGHT_COLOR)
	}
	return columnStyle.Render(components.ListView(currentDir, cursor))
}
