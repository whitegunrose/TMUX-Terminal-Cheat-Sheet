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
			{Key: "$ tmux / $ tmux new / $ tmux new-session", Desc: "Start a new session"},
			{Key: ": new", Desc: "Start a new session (tmux command mode)"},
			{Key: "$ tmux new-session -A -s mysession", Desc: "Start a new session or attach to an existing session named mysession"},
			{Key: "$ tmux new -s mysession / : new -s mysession", Desc: "Start a new session with the name mysession"},
			{Key: ": kill-session", Desc: "Kill/delete the current session"},
			{Key: "$ tmux kill-ses -t mysession / $ tmux kill-session -t mysession", Desc: "Kill/delete session mysession"},
			{Key: "$ tmux kill-session -a", Desc: "Kill/delete all sessions but the current"},
			{Key: "$ tmux kill-session -a -t mysession", Desc: "Kill/delete all sessions but mysession"},
			{Key: "Ctrl + b $", Desc: "Rename session"},
			{Key: "Ctrl + b d", Desc: "Detach from session"},
			{Key: ": attach -d", Desc: "Detach others on the session (maximize window by detaching other clients)"},
			{Key: "$ tmux ls / $ tmux list-sessions", Desc: "List all sessions"},
			{Key: "Ctrl + b s", Desc: "Show all sessions"},
			{Key: "$ tmux a / $ tmux at / $ tmux attach / $ tmux attach-session", Desc: "Attach to last session"},
			{Key: "$ tmux a -t mysession / $ tmux at -t mysession / $ tmux attach -t mysession / $ tmux attach-session -t mysession", Desc: "Attach to a session with the name mysession"},
			{Key: "Ctrl + b w", Desc: "Session and window preview"},
			{Key: "Ctrl + b (", Desc: "Move to previous session"},
			{Key: "Ctrl + b )", Desc: "Move to next session"},
		}},
	{
		Title: "Windows",
		Bindings: []Binding{
			{Key: "$ tmux new -s mysession -n mywindow", Desc: "Start a new session with the name mysession and window mywindow"},
			{Key: "Ctrl + b c", Desc: "Create window"},
			{Key: "Ctrl + b ,", Desc: "Rename current window"},
			{Key: "Ctrl + b &", Desc: "Close current window"},
			{Key: "Ctrl + b w", Desc: "List windows"},
			{Key: "Ctrl + b p", Desc: "Previous window"},
			{Key: "Ctrl + b n", Desc: "Next window"},
			{Key: "Ctrl + b 0 ... 9", Desc: "Switch/select window by number"},
			{Key: "Ctrl + b l", Desc: "Toggle last active window"},
			{Key: "Ctrl + b <", Desc: "Open window actions menu"},
			{Key: ": swap-window -s 2 -t 1", Desc: "Reorder window, swap window number 2 (src) and 1 (dst)"},
			{Key: ": swap-window -t -1", Desc: "Move current window to the left by one position"},
			{Key: ": move-window -s src_ses:win -t target_ses:win / : movew -s foo:0 -t bar:9 / : movew -s 0:0 -t 1:9", Desc: "Move window from source to target"},
			{Key: ": move-window -s src_session:src_window / : movew -s 0:9", Desc: "Reposition window in the current session"},
			{Key: ": move-window -r / : movew -r", Desc: "Renumber windows to remove gap in the sequence"},
		},
	},
	{
		Title: "Panes",
		Bindings: []Binding{
			{Key: "Ctrl + b ;", Desc: "Toggle last active pane"},
			{Key: ": split-window -h / Ctrl + b %", Desc: "Split the current pane with a vertical line to create a horizontal layout"},
			{Key: ": split-window -v / Ctrl + b \"", Desc: "Split the current pane with a horizontal line to create a vertical layout"},
			{Key: ": join-pane -s 2 -t 1", Desc: "Join two windows as panes (merge window 2 to window 1 as panes)"},
			{Key: ": join-pane -s 2.1 -t 1.0", Desc: "Move pane from one window to another (move pane 1 from window 2 to pane after 0 of window 1)"},
			{Key: "Ctrl + b {", Desc: "Move the current pane left"},
			{Key: "Ctrl + b }", Desc: "Move the current pane right"},
			{Key: "Ctrl + b ↑ / Ctrl + b ↓ / Ctrl + b → / Ctrl + b ←", Desc: "Switch to pane in the direction"},
			{Key: ": setw synchronize-panes", Desc: "Toggle synchronize-panes (send command to all panes)"},
			{Key: "Ctrl + b Spacebar", Desc: "Toggle between pane layouts"},
			{Key: "Ctrl + b o", Desc: "Switch to next pane"},
			{Key: "Ctrl + b q", Desc: "Show pane numbers"},
			{Key: "Ctrl + b q 0 ... 9", Desc: "Switch/select pane by number"},
			{Key: "Ctrl + b z", Desc: "Toggle pane zoom"},
			{Key: "Ctrl + b !", Desc: "Convert pane into a window"},
			{Key: "Ctrl + b + ↑ / Ctrl + b Ctrl + ↑ / Ctrl + b + ↓ / Ctrl + b Ctrl + ↓", Desc: "Resize current pane height (holding second key is optional)"},
			{Key: "Ctrl + b + → / Ctrl + b Ctrl + → / Ctrl + b + ← / Ctrl + b Ctrl + ←", Desc: "Resize current pane width (holding second key is optional)"},
			{Key: "Ctrl + b x", Desc: "Close current pane"},
			{Key: "Ctrl + b >", Desc: "Open pane actions menu"},
		}},
	{
		Title: "Copy Mode",
		Bindings: []Binding{
			{Key: ": setw -g mode-keys vi", Desc: "Use vi keys in buffer"},
			{Key: "Ctrl + b [", Desc: "Enter copy mode"},
			{Key: "Ctrl + b PgUp", Desc: "Enter copy mode and scroll one page up"},
			{Key: "q", Desc: "Quit mode"},
			{Key: "g", Desc: "Go to top line"},
			{Key: "G", Desc: "Go to bottom line"},
			{Key: "↑", Desc: "Scroll up"},
			{Key: "↓", Desc: "Scroll down"},
			{Key: "h", Desc: "Move cursor left"},
			{Key: "j", Desc: "Move cursor down"},
			{Key: "k", Desc: "Move cursor up"},
			{Key: "l", Desc: "Move cursor right"},
			{Key: "w", Desc: "Move cursor forward one word at a time"},
			{Key: "b", Desc: "Move cursor backward one word at a time"},
			{Key: "/", Desc: "Search forward"},
			{Key: "?", Desc: "Search backward"},
			{Key: "n", Desc: "Next keyword occurrence"},
			{Key: "N", Desc: "Previous keyword occurrence"},
			{Key: "Spacebar", Desc: "Start selection"},
			{Key: "Esc", Desc: "Clear selection"},
			{Key: "Enter", Desc: "Copy selection"},
			{Key: "Ctrl + b ]", Desc: "Paste contents of buffer_0"},
			{Key: ": show-buffer", Desc: "Display buffer_0 contents"},
			{Key: ": capture-pane", Desc: "Copy entire visible contents of pane to a buffer"},
			{Key: ": list-buffers", Desc: "Show all buffers"},
			{Key: ": choose-buffer", Desc: "Show all buffers and paste selected"},
			{Key: ": save-buffer buf.txt", Desc: "Save buffer contents to buf.txt"},
			{Key: ": delete-buffer -b 1", Desc: "Delete buffer_1"},
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
			{Key: "Ctrl + b :", Desc: "Enter command mode"},
			{Key: ": set -g OPTION", Desc: "Set OPTION for all sessions"},
			{Key: ": setw -g OPTION", Desc: "Set OPTION for all windows"},
			{Key: ": set mouse on", Desc: "Enable mouse mode"},
			{Key: "$ tmux -V", Desc: "Print tmux version"},
		},
	},
	{
		Title: "Help",
		Bindings: []Binding{
			{Key: "$ tmux list-keys / : list-keys / Ctrl + b ?", Desc: "List key bindings (shortcuts)"},
			{Key: "$ tmux info", Desc: "Show every session, window, pane, etc."},
		},
	},
}
