package ui

import "file_manager/cmd/components"

func SearchPane(currentSelectedPane int) string {
	return components.Pane("Search", "", 122, 2, currentSelectedPane == 4)
}
