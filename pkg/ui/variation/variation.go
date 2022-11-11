package variation

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

type BackMsg bool

func selectOptionCmd(variation square.Variation) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg{Variation: variation}
	}
}

type SelectMsg struct {
	Variation square.Variation
}
type Model struct {
	Item       square.CatalogItem
	Options    []square.Variation
	windowSize tea.WindowSizeMsg
	cursor     int
}

func New(item square.CatalogItem, ws tea.WindowSizeMsg) Model {
	opts, err := square.GetItemOptions(item)
	if err != nil {
		log.Fatal(err)
	}
	return Model{
		Item:       item,
		windowSize: ws,
		Options:    opts,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	if len(m.Options) < 2 {
		cmd = selectOptionCmd(m.Options[0])
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.Options)-1 {
				m.cursor++
			}
		case "enter":
			cmd = selectOptionCmd(m.Options[m.cursor])
		case "backspace":
			cmd = backSpaceCmd()
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s string
	s += fmt.Sprintf("%s \n\n", m.Item.ItemData.Name)
	s += m.RenderOptions()
	s += components.Footer()
	return s
}

func (m Model) RenderOptions() string {
	s := "--- Options ---\n\n"
	for i, opt := range m.Options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s \n", cursor, opt.ItemVariationData.Name)
	}
	return s
}
