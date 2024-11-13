package ui

import (
	"file_manager/cmd/dir"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type windowSize struct {
	width  int
	height int
}
type model struct {
	windowSize     windowSize
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
	rightNavStyle := lipgloss.NewStyle().Width(30).Border(lipgloss.NormalBorder()).Height(2)
	folderSelectStyle := lipgloss.NewStyle().Width(30).Border(lipgloss.NormalBorder()).Height(19)

	RightNav := lipgloss.JoinVertical(lipgloss.Top, rightNavStyle.Render("File Manager"), folderSelectStyle.Render("Favorites"), folderSelectStyle.Render("Recents"))
	mainSection := lipgloss.JoinVertical(lipgloss.Top, searchStyle.Render("Search"), lipgloss.JoinHorizontal(
		lipgloss.Left,
		columnStyle.Render(ListView(m.currDir, m.cursor)),
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
