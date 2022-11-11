package item

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	square "github.com/reverendheat/gccr-tui/pkg/client"
	"github.com/reverendheat/gccr-tui/pkg/ui/components"
)

func backSpaceCmd() tea.Cmd {
	return func() tea.Msg {
		return BackMsg(true)
	}
}

func selectItemCmd(catalogItem square.CatalogItem) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg{CatalogItem: catalogItem}
	}
}

type SelectMsg struct {
	CatalogItem square.CatalogItem
}

type BackMsg bool
type Model struct {
	itemList   []square.CatalogItem
	windowSize tea.WindowSizeMsg
	cursor     int
}

func New(acId string, ws tea.WindowSizeMsg) Model {
	itemList, err := square.GetItemsByCategoryId(acId)
	if err != nil {
		log.Fatal("Unable to retrieve Category items with given ID.")
	}
	return Model{
		itemList:   itemList,
		windowSize: ws,
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
			if m.cursor < len(m.itemList)-1 {
				m.cursor++
			}
		case "enter":
			cmd = selectItemCmd(m.itemList[m.cursor])
		case "backspace":
			cmd = backSpaceCmd()
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s string
	if len(m.itemList) == 0 {
		s += "Unfortunately no items in this category are available :(...\n\n"
	} else {
		s += "Good Choice...\n\n"
		for i, obj := range m.itemList {
			cursor := " "
			if i == m.cursor {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, obj.ItemData.Name)
		}
	}

	s += components.Footer()
	return s
}
