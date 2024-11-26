package ui

func SearchPane(currentSelectedPane int) string {
	return Pane("Search", "", 122, 2, currentSelectedPane == 4)
}
