package zencmd

import "github.com/agambier/zenity/internal/win"

// KillParent is internal.
func KillParent() {
	win.GenerateConsoleCtrlEvent(win.CTRL_BREAK_EVENT, 0)
}
