package menu

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	square "github.com/reverendheat/gccr-tui/pkg/client"
	"github.com/reverendheat/gccr-tui/pkg/ui/components"
)

func selectCategoryCmd(categoryId string) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg{CategoryId: categoryId}
	}
}

type SelectMsg struct {
	CategoryId string
}

type Model struct {
	categories []square.CategoryObject
	cursor     int
}

func New() Model {
	cn, err := square.GetCategoryNames()
	if err != nil {
		os.Exit(1)
	}
	return Model{
		categories: cn,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.categories)-1 {
				m.cursor++
			}
		case "enter":
			cmd = selectCategoryCmd(m.getActiveCategoryID())
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := "Categories\n\n"
	for i, category := range m.categories {
		cursorText := " "
		if i == m.cursor {
			cursorText = ">"
		}
		s += fmt.Sprintf("%s %s \n", cursorText, category.CategoryData.Name)
	}
	s += components.Footer()
	return s
}

func (m Model) getActiveCategoryID() string {
	return m.categories[m.cursor].ID
}
