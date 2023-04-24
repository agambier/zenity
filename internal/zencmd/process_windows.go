package zencmd

import "github.com/agambier/zenity/v23/internal/win"

// KillParent is internal.
func KillParent() {
	win.GenerateConsoleCtrlEvent(win.CTRL_BREAK_EVENT, 0)
}
