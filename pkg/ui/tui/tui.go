package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	square "github.com/reverendheat/gccr-tui/pkg/client"
	"github.com/reverendheat/gccr-tui/pkg/ui/item"
	"github.com/reverendheat/gccr-tui/pkg/ui/menu"
	"github.com/reverendheat/gccr-tui/pkg/ui/modifier"
	"github.com/reverendheat/gccr-tui/pkg/ui/variation"
)

type sessionState struct {
	activeScreen      activeScreen
	currentCategoryId string
	currentItem       square.CatalogItem
	currentVariation  square.Variation
}

func newSessionState() sessionState {
	return sessionState{
		activeScreen: menuView,
	}
}

type activeScreen int

const (
	menuView activeScreen = iota
	itemView
	variationView
	modifierView
)

type MainModel struct {
	state      sessionState
	menu       tea.Model
	item       tea.Model
	variation  tea.Model
	modifier   tea.Model
	windowSize tea.WindowSizeMsg
}

func New() MainModel {
	return MainModel{
		state: newSessionState(),
		menu:  menu.New(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg // pass this along to the entry view so it uses the full window size when it's initialized
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case menu.SelectMsg:
		m.state.currentCategoryId = msg.CategoryId
		m.item = item.New(m.state.currentCategoryId, m.windowSize)
		m.state.activeScreen = itemView
	case item.BackMsg:
		m.state.activeScreen = menuView
	case item.SelectMsg:
		m.state.currentItem = msg.CatalogItem
		m.variation = variation.New(m.state.currentItem, m.windowSize)
		m.state.activeScreen = variationView
	case variation.BackMsg:
		m.state.activeScreen = itemView
	case variation.SelectMsg:
		m.state.currentVariation = msg.Variation
		m.modifier = modifier.New(m.state.currentItem, m.windowSize)
		m.state.activeScreen = modifierView
	case modifier.BackMsg:
		m.state.activeScreen = variationView
	}

	switch m.state.activeScreen {
	case menuView:
		newMenu, newCmd := m.menu.Update(msg)
		menuModel, ok := newMenu.(menu.Model)
		if !ok {
			panic("could not perform assertion on menu model")
		}
		m.menu = menuModel
		cmd = newCmd
	case itemView:
		newItem, newCmd := m.item.Update(msg)
		itemModel, ok := newItem.(item.Model)
		if !ok {
			panic("could not perform assertion on item model")
		}
		m.item = itemModel
		cmd = newCmd
	case variationView:
		newvariation, newCmd := m.variation.Update(msg)
		variationModel, ok := newvariation.(variation.Model)
		if !ok {
			panic("could not perform assertion on variation model")
		}
		m.variation = variationModel
		cmd = newCmd
	case modifierView:
		newModifier, newCmd := m.modifier.Update(msg)
		modifierModel, ok := newModifier.(modifier.Model)
		if !ok {
			panic("could not perform assertion on modifier model")
		}
		m.modifier = modifierModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View return the text UI to be output to the terminal
func (m MainModel) View() string {

	switch m.state.activeScreen {
	case itemView:
		return m.item.View()
	case variationView:
		return m.variation.View()
	case modifierView:
		return m.modifier.View()
	default:
		return m.menu.View()
	}
}

func Start() {
	m := New()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Sorry friend, there's been an error: %v", err)
		os.Exit(1)
	}
}
