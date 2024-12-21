package ui

import (
	"file_manager/cmd/components"
	"file_manager/cmd/types"
)

func RecentsPane(currentSelectedPane int) string {
	return components.Pane("Recents", RenderRecentsPaneList(), 30, 19, currentSelectedPane == 3)
}

func RenderRecentsPaneList() string {
	items := []types.Item{}
	cursor := 0
	return components.ListView(items, cursor)
}
