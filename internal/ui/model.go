package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/whitegunrose/TMUX-Terminal-Cheat-Sheet/internal/data"
)

// Model holds the entire application state.
// In bubbletea, all state lives here — nothing is mutated outside of Update().
type Model struct {
	cursor int // which category is selected
	width  int
	height int
}

// New returns the initial model.
func New() Model {
	return Model{}
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

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		// case "up", "k":
		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}

		// case "down", "j":
		case "right", "l":
			if m.cursor < len(data.Sections)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

// View renders the UI from the current model state.
// This is a pure function — it never modifies the model.
func (m Model) View() string {
	if m.width == 0 {
		return "" // not yet sized; avoid a blank flash
	}

	title := titleStyle.Render("󰀲  TMUX Cheat Sheet")

	// left := m.renderCategories()
	top := m.renderCategories()
	// right := m.renderBindings()
	bottom := m.renderBindings()

	// panels := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	panels := lipgloss.JoinVertical(lipgloss.Top, top, bottom)
	footer := helpStyle.Render("↑/↓ or j/k to navigate  •  q to quit")

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
			// b.WriteString(selectedCategoryStyle.Render("▶ "+label) + "\n")
			b.WriteString(selectedCategoryStyle.Render("▶ "+label) + " ")
		} else {
			// b.WriteString(categoryStyle.Render("  "+label) + "\n")
			b.WriteString(categoryStyle.Render("  "+label) + " ")
		}
	}

	// return leftPanelStyle.Render(b.String())
	return topPanelStyle.Render(b.String())
}

// renderBindings builds the right panel showing keybindings for the
// currently selected section.
func (m Model) renderBindings() string {
	section := data.Sections[m.cursor]

	var b strings.Builder
	b.WriteString(sectionTitleStyle.Render(section.Title) + "\n\n")

	for _, binding := range section.Bindings {
		key := keyStyle.Render(binding.Key)
		desc := descStyle.Render(binding.Desc)
		b.WriteString(fmt.Sprintf("%s %s\n", key, desc))
	}

	// return rightPanelStyle.Render(b.String())
	return bottomPanelStyle.Render(b.String())
}
