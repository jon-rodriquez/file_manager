package recents

import (
	"file_manager/cmd/components"
	"file_manager/cmd/types"

	tea "github.com/charmbracelet/bubbletea"
)

type RecentsPane struct {
	name                   string
	width                  int
	height                 int
	location               int
	cursor                 int
	recentsList            []types.Item
	childrenOfSelectedItem []types.Item
}

func NewRecentsPane() *RecentsPane {
	pane := &RecentsPane{}
	pane.Initialize()
	return pane
}

func (pane *RecentsPane) Initialize() {
	pane.name = "Recents"
	pane.width = 30
	pane.height = 19
	pane.location = 3
	pane.cursor = 0
	pane.childrenOfSelectedItem = []types.Item{}
	pane.recentsList = []types.Item{}
}

func (pane *RecentsPane) Update(msg tea.Msg) {
}

func (pane *RecentsPane) RenderPane(selectedPane int) string {
	return components.Pane(pane.name, "", pane.width, pane.height, selectedPane == pane.location)
}
