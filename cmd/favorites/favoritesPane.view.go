package favorites

import (
	"file_manager/cmd/components"
	"file_manager/cmd/types"

	tea "github.com/charmbracelet/bubbletea"
)

type FavoritesPane struct {
	cursor                 int
	favoritesList          []types.Item
	childrenOfSelectedItem []types.Item
}

func NewFavoritesPane() *FavoritesPane {
	pane := &FavoritesPane{}
	pane.Initialize()
	return pane
}

func (pane *FavoritesPane) Initialize() {
	pane.favoritesList = GetFavorties()
	pane.cursor = 0
	pane.childrenOfSelectedItem = []types.Item{}
}

func (pane *FavoritesPane) Update(msg tea.Msg) (*FavoritesPane, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if pane.cursor < len(pane.favoritesList)-1 {
				pane.cursor++
			}
		case "k":
			if pane.cursor > 0 {
				pane.cursor--
			}
		}
	}
	return pane, nil
}

func (pane *FavoritesPane) RenderPane(currentSelectedPane int) string {
	// Print the styled content with the title
	return components.Pane("Favorites", pane.renderFavoritesPaneList(), 30, 19, currentSelectedPane == 2)
}

// Print the styled content with a title
func (pane *FavoritesPane) renderFavoritesPaneList() string {
	if pane.favoritesList == nil {
		pane.favoritesList = GetFavorties()
	}
	cursor := 0
	return components.ListView(pane.favoritesList, cursor)
}
