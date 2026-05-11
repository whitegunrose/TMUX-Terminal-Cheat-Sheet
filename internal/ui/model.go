package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/whitegunrose/TMUX-Terminal-Cheat-Sheet/internal/data"
)

// Model holds the entire application state.
// In bubbletea, all state lives here — nothing is mutated outside of Update().
type Model struct {
	cursor   int // which category is selected
	width    int
	height   int
	viewport viewport.Model
}

// New returns the initial model.
func New() Model {
	vp := viewport.New(0, 0)
	vp.SetContent("")
	return Model{viewport: vp}
}

// Init runs any startup commands. We don't need any, so return nil.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles all incoming messages (key presses, window resizes, etc.)
// and returns an updated copy of the model plus an optional command.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 3 // title + padding
		footerHeight := 2 // help text + padding

		m.viewport.Width = m.width - 6
		m.viewport.Height = m.height - headerHeight - footerHeight

		// re-render content whenever the pane is resized
		m.viewport.SetContent(m.renderBindings())

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
				m.viewport.SetContent(m.renderBindings())
				m.viewport.GotoTop() // reset scroll on category change
			}

		case "right", "l":
			if m.cursor < len(data.Sections)-1 {
				m.cursor++
				m.viewport.SetContent(m.renderBindings())
				m.viewport.GotoTop()
			}
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// View renders the UI from the current model state.
// This is a pure function — it never modifies the model.
func (m Model) View() string {
	if m.width == 0 {
		return "" // not yet sized; avoid a blank flash
	}

	title := titleStyle.Render("󰀲  TMUX Cheat Sheet")

	top := m.renderCategories()
	bottom := bottomPanelStyle.Render(m.viewport.View())

	panels := lipgloss.JoinVertical(lipgloss.Top, top, bottom)
	footer := helpStyle.Render("← / → or h/l to navigate  •  q to quit")

	return appStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, title, panels, footer),
	)
}

// renderCategories builds the left panel listing all section titles.
func (m Model) renderCategories() string {
	var b strings.Builder

	for i, section := range data.Sections {
		label := section.Title
		if i == m.cursor {
			b.WriteString(selectedCategoryStyle.Render("▶ "+label) + " ")
		} else {
			b.WriteString(categoryStyle.Render("  "+label) + " ")
		}
	}

	// return leftPanelStyle.Render(b.String())
	// return topPanelStyle.Render(b.String())
	return topPanelStyle.Width(m.width - 6).Render(b.String()) // dynamic width
}

// renderBindings builds the right panel showing keybindings for the
// currently selected section.
func (m Model) renderBindings() string {
	section := data.Sections[m.cursor]

	keyWidth := m.viewport.Width / 2
	descWidth := m.viewport.Width - 2 // total width minus spacing

	var b strings.Builder
	b.WriteString(sectionTitleStyle.Render(section.Title) + "\n\n")

	for _, binding := range section.Bindings {
		key := keyStyle.Width(keyWidth).Render(binding.Key)
		desc := descStyle.Width(descWidth).Render(binding.Desc) // dynamic width
		b.WriteString(fmt.Sprintf("%s\n%s\n\n", key, desc))
	}

	return b.String()
}
