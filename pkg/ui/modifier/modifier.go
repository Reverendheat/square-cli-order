package modifier

import (
	"errors"
	"fmt"
	"log"
	"reflect"

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
	selections    map[string]interface{}
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
		selections:    initSelections(mods),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func contains(l []int, n int) bool {
	for _, v := range l {
		if v == n {
			return true
		}
	}
	return false
}

func initSelections(mods []square.ModifierList) map[string]interface{} {
	selections := make(map[string]interface{})
	for _, ml := range mods {
		if ml.ModifierListData.SelectionType == "SINGLE" {
			selections[ml.ID] = nil
		} else {
			selections[ml.ID] = []int{}
		}
	}
	return selections
}

func findIndexByValue(s []int, val int) (int, error) {
	for i, v := range s {
		if val == v {
			return i, nil
		}
	}
	return val, errors.New("Value you not found in slice")
}

func removeAtIndex(s []int, idx int) []int {
	s[idx] = s[0]
	return s[1:]
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
		case "enter":
			//IT HAS TO BE INDIVIDUAL MODIFIER NAME ON THE SELECTION NOT THE ID FOR THE WHOLE LIST!
			if val, ok := m.selections[m.activeMl.ID]; ok {
				//Modifier list exists...
				if m.activeMl.ModifierListData.SelectionType == "SINGLE" {
					// single..
					if val == m.cursor {
						// empty it
						m.selections[m.activeMl.ID] = nil
					} else {
						// add it
						m.selections[m.activeMl.ID] = m.cursor
					}
				} else if m.activeMl.ModifierListData.SelectionType == "MULTIPLE" {
					tmpSlice := reflect.ValueOf(m.selections[m.activeMl.ID]).Interface().([]int)
					if contains(tmpSlice, m.cursor) {
						idx, err := findIndexByValue(tmpSlice, m.cursor)
						if err != nil {
							log.Fatalf("%d was not found in activeMl list", m.cursor)
						}
						m.selections[m.activeMl.ID] = removeAtIndex(tmpSlice, idx)
					} else {
						m.selections[m.activeMl.ID] = append(tmpSlice, m.cursor)
					}
				}
			} else {
				// Modifier list key does not exist
				if m.activeMl.ModifierListData.SelectionType == "SINGLE" {
					m.selections[m.activeMl.ID] = m.cursor
				}
				if m.activeMl.ModifierListData.SelectionType == "MULTIPLE" {
					tmpSlice := []int{m.cursor}
					m.selections[m.activeMl.ID] = tmpSlice
				}
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

func (m Model) RenderModifiers(ml square.ModifierListData, id string) string {
	var s string
	for i, mod := range ml.Modifiers {
		cursor := " "
		selected := " "
		if m.cursor == i && ml.Name == m.activeMl.ModifierListData.Name {
			cursor = ">"
		}
		if ml.SelectionType == "SINGLE" {
			if m.selections[id] == i {
				selected = "X"
			}
		} else if ml.SelectionType == "MULTIPLE" {
			tmpSlice := reflect.ValueOf(m.selections[id]).Interface().([]int)
			if contains(tmpSlice, i) {
				selected = "X"
			}
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, mod.ModifierData.Name)
	}
	return s
}

func (m Model) RenderModifierLists() string {
	s := "--- Mods ---\n\n"
	for _, ml := range m.Modifiers {
		s += fmt.Sprintf("-- %s -- ", ml.ModifierListData.Name)
		if ml.ModifierListData.SelectionType == "SINGLE" {
			s += "Choose one\n"
		} else {
			s += "Choose many\n"
		}
		s += m.RenderModifiers(ml.ModifierListData, ml.ID)
	}
	return s
}
