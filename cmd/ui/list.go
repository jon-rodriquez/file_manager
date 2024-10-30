package ui

import (
	"file_manager/cmd/dir"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	name  string
	path  string
	isDir bool
}

type model struct {
	currDir []Item
	cursor  int
}

func InitialModel() model {
	dirs := dir.ListDir()
	currentDirItems := []Item{}

	for _, item := range dirs {
		currentDirItems = append(currentDirItems, Item{
			name:  item.Name(),
			path:  item.Name(),
			isDir: item.IsDir(),
		})
	}

	return model{
		currDir: currentDirItems,
		cursor:  0,
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
			}
		case "k":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "current dir\n"
	selectedStyle := lipgloss.NewStyle().Background(lipgloss.Color("202"))
	for i, item := range m.currDir {
		var folder string

		if item.isDir == true {
			folder = "📁"
		} else {
			folder = "📄"
		}

		if m.cursor == i {

			temp := fmt.Sprintf("[%s] %s\n", folder, item.name)
			s += selectedStyle.Render(temp)
		} else {
			s += fmt.Sprintf("[%s] %s\n", folder, item.name)
		}

	}
	s += "\nPress q to quit.\n"

	return s
}
