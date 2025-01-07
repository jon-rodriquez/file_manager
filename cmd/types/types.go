package types

import tea "github.com/charmbracelet/bubbletea"

type Item struct {
	Name  string
	Path  string
	IsDir bool
}

type Initializer interface {
	Initialize()
}

type PaneRenderer interface {
	RenderPane(curr int) string
}

type Updater interface {
	Update(msg tea.Msg)
}

type MainPaneRenderer interface {
	RenderMainPane() string
}

type Pane interface {
	Initializer
	PaneRenderer
	Updater
	MainPaneRenderer
}
