package ui

import (
	"file_manager/cmd/dir"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
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
	}
	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		ListView(m.currDir, m.cursor),
		ListView(m.childDir, -1))
}

func (m *model) updateChildDir() {
	if m.currDir[m.cursor].IsDir {
		m.childDir = dir.GetDirItems(m.currentDirPath + "/" + m.currDir[m.cursor].Name)
	} else {
		m.childDir = []dir.Item{}
	}
}
