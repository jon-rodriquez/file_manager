package favorites

import (
	"file_manager/cmd/components"
)

func FavoritesPane(currentSelectedPane int) string {
	// Print the styled content with the title
	return components.Pane("Favorites", RenderFavoritesPaneList(), 30, 19, currentSelectedPane == 2)
}

// Print the styled content with a title

func RenderFavoritesPaneList() string {
	items := GetFavorties()
	cursor := 0
	return components.ListView(items, cursor)
}
