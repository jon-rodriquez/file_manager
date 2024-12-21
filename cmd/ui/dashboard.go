package ui

import (
	"file_manager/cmd/favorites"
	fileManagerPane "file_manager/cmd/fileManager"
	"file_manager/cmd/recents"
	"file_manager/cmd/types"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var HIGHLIGHT_COLOR = lipgloss.Color("#ff7f36")

type windowSize struct {
	width  int
	height int
}
type model struct {
	windowSize   windowSize
	panes        []types.Pane
	selectedPane int
	entryPath    string
}

func InitialModel() model {
	currentDirPath, _ := os.Getwd()

	return model{
		entryPath:    currentDirPath,
		selectedPane: 1,
		panes: []types.Pane{
			fileManagerPane.NewFileManagerPane(),
			favorites.NewFavoritesPane(),
			recents.NewRecentsPane(),
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	selected := m.panes[m.selectedPane-1]

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j":
		case "k":
		case "h":
		case "l":
			selected.Update(msg)
		case "1":
			m.selectedPane = 1
		case "2":
			m.selectedPane = 2
		case "3":
			m.selectedPane = 3
		case "4":
			m.selectedPane = 4

		}
	case tea.WindowSizeMsg:
		m.windowSize.width = msg.Width
		m.windowSize.height = msg.Height

	}
	return m, nil
}

func (m model) View() string {
	// columnStyle := lipgloss.NewStyle().Width(60).Border(lipgloss.NormalBorder()).Height(40)

	RightNav := lipgloss.JoinVertical(lipgloss.Top, m.panes[0].RenderPane(m.selectedPane), m.panes[1].RenderPane(m.selectedPane), m.panes[2].RenderPane(m.selectedPane))
	mainSection := lipgloss.JoinVertical(lipgloss.Top, SearchPane(m.selectedPane))

	return lipgloss.JoinHorizontal(lipgloss.Left, RightNav, mainSection)
}
