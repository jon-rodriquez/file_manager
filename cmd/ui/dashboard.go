package ui

import (
	"file_manager/cmd/dir"
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
	selectedPane int
	// maybe an array of panes.
	// those panes will have the render function for the specific pane but also a funciton that returns main view section.

	currDir        []dir.Item
	childDir       []dir.Item
	currentDirPath string
	cursor         int
}

func InitialModel() model {
	currentDirPath, _ := os.Getwd()
	currentDirItems := dir.GetDirItems(currentDirPath)
	childDirItems := []dir.Item{}

	if currentDirItems[0].IsDir {
		childDirItems = dir.GetDirItems(currentDirPath + "/" + currentDirItems[0].Name)
	}

	return model{
		currDir:        currentDirItems,
		childDir:       childDirItems,
		currentDirPath: currentDirPath,
		cursor:         0,
		selectedPane:   5,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j":
			if m.cursor < len(m.currDir)-1 {
				m.cursor++
				m.updateChildDir()
			}
		case "k":
			if m.cursor > 0 {
				m.cursor--
				m.updateChildDir()
			}
		case "h":
			// TODO: need to handle the case when the current directory is root directory
			// currently throws an error
			parentDir, dirs := dir.MoveToParentDir(m.currentDirPath)
			m.currentDirPath = parentDir
			m.currDir = dirs
			m.cursor = 0
			m.updateChildDir()
		case "l":
			if m.currDir[m.cursor].IsDir {
				m.currentDirPath = m.currentDirPath + "/" + m.currDir[m.cursor].Name
				m.currDir = dir.GetDirItems(m.currentDirPath)
				m.cursor = 0
				m.updateChildDir()
			}
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
	columnStyle := lipgloss.NewStyle().Width(60).Border(lipgloss.NormalBorder()).Height(40)
	searchStyle := lipgloss.NewStyle().Width(122).Border(lipgloss.NormalBorder()).Height(2)

	RightNav := lipgloss.JoinVertical(lipgloss.Top, FileManagerPane(m.selectedPane), FavoritesPane(m.selectedPane), RecentsPane(m.selectedPane))
	mainSection := lipgloss.JoinVertical(lipgloss.Top, searchStyle.Render("Search[4]"), lipgloss.JoinHorizontal(
		lipgloss.Left,
		ParentDirPane(m.selectedPane, m.currDir, m.cursor),
		columnStyle.Render(ListView(m.childDir, -1))),
	)
	return lipgloss.JoinHorizontal(lipgloss.Left, RightNav, mainSection)
}

func (m *model) updateChildDir() {
	if m.currDir[m.cursor].IsDir {
		m.childDir = dir.GetDirItems(m.currentDirPath + "/" + m.currDir[m.cursor].Name)
	} else {
		m.childDir = []dir.Item{}
	}
}
