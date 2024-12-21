package fileManagerPane

import (
	"file_manager/cmd/components"

	tea "github.com/charmbracelet/bubbletea"
)

type FileManagerPane struct {
	name     string
	width    int
	height   int
	location int
}

func NewFileManagerPane() *FileManagerPane {
	pane := &FileManagerPane{}
	pane.Initialize()
	return pane
}

func (pane *FileManagerPane) Initialize() {
	pane.name = "File Manager"
	pane.width = 30
	pane.height = 2
	pane.location = 1
}

func (pane *FileManagerPane) Update(msg tea.Msg) {
	// No update logic for this pane
}

func (pane *FileManagerPane) RenderPane(selectedPane int) string {
	return components.Pane(pane.name, "", pane.width, pane.height, selectedPane == pane.location)
}
