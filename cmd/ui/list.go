package ui

import (
	"file_manager/cmd/dir"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

type Item struct {
	name  string
	path  string
	isDir bool
}

type model struct {
	currDir        []Item
	childDir       []Item
	currentDirPath string
	cursor         int
}

func InitialModel() model {
	currentDirPath, _ := os.Getwd()
	dirs := dir.ListDir(currentDirPath)
	currentDirItems := []Item{}
	childDirItems := []Item{}

	for _, item := range dirs {
		currentDirItems = append(currentDirItems, Item{
			name:  item.Name(),
			path:  item.Name(),
			isDir: item.IsDir(),
		})
		childDirItems = append(currentDirItems, Item{
			name:  item.Name(),
			path:  item.Name(),
			isDir: item.IsDir(),
		})
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
			}
		case "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "h":
      //TODO: need to handle the case when the current directory is root directory
			parentDir, dirs := dir.MoveToParentDir(m.currentDirPath)
			m.currentDirPath = parentDir
			currentDirItems := []Item{}
			for _, item := range dirs {
				currentDirItems = append(currentDirItems, Item{
					name:  item.Name(),
					path:  item.Name(),
					isDir: item.IsDir(),
				})
				m.currDir = currentDirItems
			}
			m.cursor = 0
    case "l":
      
		}
	}
	return m, nil
}

func (m model) View() string {
	baseStyle := lipgloss.NewStyle()

	dimColor := lipgloss.Color("250")
	hightlightColor := lipgloss.Color("#EE6FF8")

	selected := m.cursor

	l := list.New().Enumerator(func(_ list.Items, i int) string {
		if i == selected {
			return " "
		}
		return " "
	}).
		ItemStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			st := baseStyle
			if selected == i {
				return st.Foreground(hightlightColor).Background(lipgloss.Color("236"))
			}
			return st.Foreground(dimColor)
		}).
		EnumeratorStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			if selected == i {
				return lipgloss.NewStyle().Foreground(hightlightColor)
			}
			return lipgloss.NewStyle().Foreground(dimColor)
		})

	for _, item := range m.currDir {

		var folder string

		if item.isDir == true {
			folder = "\uf07c"
		} else {
			folder = "\uf15b"
		}
		l.Item(folder + " " + item.name)
	}

	s := fmt.Sprintf("%s", l)

	return s
}
