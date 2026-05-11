package data

// Binding represents a single TMUX keybinding or command.
type Binding struct {
	Key  string
	Desc string
}

// Section groups related bindings under a title.
type Section struct {
	Title    string
	Bindings []Binding
}

// Sections is the full cheat sheet content. Add or edit freely.
var Sections = []Section{
	{
		Title: "Sessions",
		Bindings: []Binding{
			{Key: "tmux new -s <name>", Desc: "Start a new named session"},
			{Key: "tmux ls", Desc: "List all sessions"},
			{Key: "tmux attach -t <name>", Desc: "Attach to a named session"},
			{Key: "tmux kill-session -t <name>", Desc: "Kill a named session"},
			{Key: "prefix + $", Desc: "Rename the current session"},
			{Key: "prefix + d", Desc: "Detach from the current session"},
			{Key: "prefix + s", Desc: "Interactively switch sessions"},
			{Key: "prefix + (", Desc: "Switch to the previous session"},
			{Key: "prefix + )", Desc: "Switch to the next session"},
		},
	},
	{
		Title: "Windows",
		Bindings: []Binding{
			{Key: "prefix + c", Desc: "Create a new window"},
			{Key: "prefix + ,", Desc: "Rename the current window"},
			{Key: "prefix + w", Desc: "List all windows"},
			{Key: "prefix + n", Desc: "Move to the next window"},
			{Key: "prefix + p", Desc: "Move to the previous window"},
			{Key: "prefix + 0–9", Desc: "Switch to window by number"},
			{Key: "prefix + &", Desc: "Kill the current window"},
			{Key: "prefix + .", Desc: "Move window to a new index"},
		},
	},
	{
		Title: "Panes",
		Bindings: []Binding{
			{Key: `prefix + %`, Desc: "Split pane vertically (side by side)"},
			{Key: `prefix + "`, Desc: "Split pane horizontally (top/bottom)"},
			{Key: "prefix + arrow", Desc: "Move focus between panes"},
			{Key: "prefix + o", Desc: "Cycle through panes"},
			{Key: "prefix + z", Desc: "Toggle zoom on the current pane"},
			{Key: "prefix + x", Desc: "Kill the current pane"},
			{Key: "prefix + {", Desc: "Swap pane with the one above"},
			{Key: "prefix + }", Desc: "Swap pane with the one below"},
			{Key: "prefix + Space", Desc: "Cycle through pane layouts"},
			{Key: "prefix + q", Desc: "Show pane numbers briefly"},
		},
	},
	{
		Title: "Copy Mode",
		Bindings: []Binding{
			{Key: "prefix + [", Desc: "Enter copy mode"},
			{Key: "q / Esc", Desc: "Exit copy mode"},
			{Key: "Space", Desc: "Start selection (vi mode)"},
			{Key: "Enter", Desc: "Copy selection and exit"},
			{Key: "prefix + ]", Desc: "Paste the most recent buffer"},
			{Key: "/ or ?", Desc: "Search forward / backward"},
			{Key: "n / N", Desc: "Next / previous search match"},
			{Key: "g / G", Desc: "Jump to top / bottom"},
		},
	},
	{
		Title: "Misc",
		Bindings: []Binding{
			{Key: "prefix + ?", Desc: "Show all keybindings"},
			{Key: "prefix + :", Desc: "Open the tmux command prompt"},
			{Key: "prefix + t", Desc: "Show a clock in the current pane"},
			{Key: "prefix + ~", Desc: "Show previous messages"},
			{Key: "tmux source ~/.tmux.conf", Desc: "Reload your config file"},
		},
	},
}
