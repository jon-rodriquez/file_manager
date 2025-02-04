package components

import (
	"file_manager/cmd/types"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func ListView(items []types.Item, cursor int) string {
	baseStyle := lipgloss.NewStyle()
	dimColor := lipgloss.Color("250")
	hightlightColor := lipgloss.Color("#ff7f36")

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
				return st.Foreground(hightlightColor).Background(lipgloss.Color("236")).Width(29)
			}
			return st.Foreground(dimColor)
		}).
		EnumeratorStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			if selected == i {
				return lipgloss.NewStyle().Foreground(hightlightColor)
			}
			return lipgloss.NewStyle().Foreground(dimColor)
		})

	limit := 20

	if len(items) < 20 {
		limit = len(items)
	}

	for i := 0; i < limit; i++ {
		item := items[i]

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
