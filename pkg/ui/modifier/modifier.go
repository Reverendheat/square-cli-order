package modifier

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

type Model struct {
	Item          square.CatalogItem
	cursor        int
	windowSize    tea.WindowSizeMsg
	activeMl      square.ModifierList
	activeMlIndex int
	activeMlMap   map[int]square.ModifierList
	selections    map[string]struct{}
	Modifiers     []square.ModifierList
}

func New(item square.CatalogItem, ws tea.WindowSizeMsg) Model {
	var activeMl square.ModifierList
	activeMlMap := make(map[int]square.ModifierList)
	mods, err := square.GetItemModifiers(item)
	if err != nil {
		log.Fatal(err)
	}
	if len(mods) > 0 {
		activeMl = mods[0]
		activeMlMap[0] = mods[0]
	}
	return Model{
		Item:          item,
		windowSize:    ws,
		Modifiers:     mods,
		activeMl:      activeMl,
		activeMlIndex: 0,
		activeMlMap:   activeMlMap,
		selections:    make(map[string]struct{}),
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
			} else if m.activeMlIndex > 0 {
				m.activeMlIndex--
				m.activeMl = m.Modifiers[m.activeMlIndex]
				m.cursor = len(m.activeMl.ModifierListData.Modifiers) - 1
			}
		case "down", "j":
			if m.cursor < len(m.activeMl.ModifierListData.Modifiers)-1 {
				m.cursor++
			} else if m.activeMlIndex < len(m.Modifiers)-1 {
				m.activeMlIndex++
				m.activeMl = m.Modifiers[m.activeMlIndex]
				m.cursor = 0
			}
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
	s += m.RenderModifierLists()
	s += components.Footer()
	return s
}

func (m Model) RenderModifiers(ml square.ModifierListData) string {
	var s string
	for i, mod := range ml.Modifiers {
		cursor := " "
		selected := " "
		if m.cursor == i && ml.Name == m.activeMl.ModifierListData.Name {
			cursor = ">"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, mod.ModifierData.Name)
	}
	return s
}

func (m Model) RenderModifierLists() string {
	s := "--- Mods ---\n\n"
	for _, ml := range m.Modifiers {
		s += fmt.Sprintf("-- %s --\n", ml.ModifierListData.Name)
		s += m.RenderModifiers(ml.ModifierListData)
	}
	return s
}
