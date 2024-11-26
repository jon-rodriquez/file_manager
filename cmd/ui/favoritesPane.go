package ui

import (
	"file_manager/cmd/dir"
)

func FavoritesPane(currentSelectedPane int) string {
	// Print the styled content with the title
	return Pane("Favorites", RenderFavoritesPaneList(), 30, 19, currentSelectedPane == 2)
}

// Print the styled content with a title

func RenderFavoritesPaneList() string {
	items := dir.GetFavorties()
	cursor := 0
	return ListView(items, cursor)
}
