package types

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

type Pane interface {
	Initializer
	PaneRenderer
}
