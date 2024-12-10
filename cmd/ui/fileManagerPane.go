package ui

import "file_manager/cmd/components"

func FileManagerPane(currentSelectedPane int) string {
	return components.Pane("File Manager", "", 30, 2, currentSelectedPane == 1)
}
