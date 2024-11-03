package ui

import (
	"file_manager/cmd/dir"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func ListView(items []dir.Item, cursor int) string {
	baseStyle := lipgloss.NewStyle()
	dimColor := lipgloss.Color("250")
	hightlightColor := lipgloss.Color("#EE6FF8")

	selected := cursor

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

	for _, item := range items {

		var folder string

		if item.IsDir == true {
			folder = "\uf07c"
		} else {
			folder = "\uf15b"
		}
		l.Item(folder + " " + item.Name)
	}

	s := fmt.Sprintf("%s", l)
	return s
}
