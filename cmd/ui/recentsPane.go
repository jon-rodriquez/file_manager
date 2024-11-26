package ui

import (
	"file_manager/cmd/dir"
)

func RecentsPane(currentSelectedPane int) string {
	return Pane("Recents", RenderRecentsPaneList(), 30, 19, currentSelectedPane == 3)
}

func RenderRecentsPaneList() string {
	items := []dir.Item{}
	cursor := 0
	return ListView(items, cursor)
}
