package ui

import "github.com/charmbracelet/lipgloss"

const (
	// leftPanelWidth = 22
	borderColor = "#3C3C3C"
	green       = "#25A065"
	orange      = "#F99417"
	blue        = "#5C7AEA"
	subtle      = "#626262"
	bright      = "#FFFDF5"
)

var (
	// Overall app padding
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	// Top panel: category list
	topPanelStyle = lipgloss.NewStyle().
		// Width(leftPanelWidth).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(borderColor)).
		BorderLeft(true).
		BorderRight(true).
		PaddingRight(1)

	// Bottom panel: bindings
	bottomPanelStyle = lipgloss.NewStyle().
				PaddingLeft(2)

	// Title bar
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(bright)).
			Background(lipgloss.Color(green)).
			Padding(0, 1).
			MarginBottom(1).
			Bold(true)

	// Category items in left panel
	categoryStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(subtle)).
			PaddingLeft(1)

	selectedCategoryStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(green)).
				PaddingLeft(1).
				Bold(true)

	// Section heading in right panel
	sectionTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(orange)).
				Bold(true).
				MarginBottom(1).
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).
				BorderForeground(lipgloss.Color(borderColor))

	// Keybinding column
	keyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(blue)).
			Width(32)

	// Description column
	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(bright))

	// Footer help text
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(subtle)).
			MarginTop(1)
)
