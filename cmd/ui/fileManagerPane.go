package ui

func FileManagerPane(currentSelectedPane int) string {
	return Pane("File Manager", "", 30, 2, currentSelectedPane == 1)
}
