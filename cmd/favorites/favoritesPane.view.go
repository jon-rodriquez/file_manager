package favorites

import (
	"file_manager/cmd/components"
	"file_manager/cmd/dir"
	"file_manager/cmd/types"

	tea "github.com/charmbracelet/bubbletea"
)

type FavoritesPane struct {
	name                   string
	width                  int
	height                 int
	location               int
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
	pane.name = "Favorites"
	pane.width = 30
	pane.height = 19
	pane.location = 2
	pane.favoritesList = GetFavorties()
	pane.cursor = 0
	pane.childrenOfSelectedItem = []types.Item{}
}

func (pane *FavoritesPane) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if pane.cursor < len(pane.favoritesList)-1 {
				pane.cursor++
				pane.childrenOfSelectedItem = dir.GetDirItems(pane.favoritesList[pane.cursor].Path)
			}
		case "k":
			if pane.cursor > 0 {
				pane.cursor--
				pane.childrenOfSelectedItem = dir.GetDirItems(pane.favoritesList[pane.cursor].Path)
			}
		}
	}
}

func (pane *FavoritesPane) RenderPane(selectedPane int) string {
	// Print the styled content with the title
	return components.Pane(pane.name, pane.renderFavoritesPaneList(), pane.width, pane.height, selectedPane == pane.location)
}

// Print the styled content with a title
func (pane *FavoritesPane) renderFavoritesPaneList() string {
	if pane.favoritesList == nil {
		pane.favoritesList = GetFavorties()
	}
	return components.ListView(pane.favoritesList, pane.cursor)
}

func (pane *FavoritesPane) RenderMainPane() string {
	return components.Pane(pane.favoritesList[pane.cursor].Name, components.ListView(pane.childrenOfSelectedItem, -1), 60, 30, true)
}
