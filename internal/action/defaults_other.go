//go:build !darwin
// +build !darwin

package action

var bufdefaults = map[string]string{
	"Up":             "CursorUp",
	"Down":           "CursorDown",
	"Right":          "CursorRight",
	"Left":           "CursorLeft",
	"ShiftUp":        "SelectUp",
	"ShiftDown":      "SelectDown",
	"ShiftLeft":      "SelectLeft",
	"ShiftRight":     "SelectRight",
	"CtrlLeft":       "WordLeft",
	"CtrlRight":      "WordRight",
	"AltUp":          "MoveLinesUp",
	"AltDown":        "MoveLinesDown",
	"CtrlShiftRight": "SelectWordRight",
	"CtrlShiftLeft":  "SelectWordLeft",
	"AltLeft":        "StartOfTextToggle",
	"AltRight":       "EndOfLine",
	"AltShiftLeft":   "SelectToStartOfTextToggle",
	"ShiftHome":      "SelectToStartOfTextToggle",
	"AltShiftRight":  "SelectToEndOfLine",
	"ShiftEnd":       "SelectToEndOfLine",
	"CtrlUp":         "CursorStart",
	"CtrlDown":       "CursorEnd",
	"CtrlShiftUp":    "SelectToStart",
	"CtrlShiftDown":  "SelectToEnd",
	"Alt-{":          "ParagraphPrevious",
	"Alt-}":          "ParagraphNext",
	"Enter":          "InsertNewline",
	"CtrlH":          "Backspace",
	"Backspace":      "Backspace",
	"OldBackspace":   "Backspace",
	"Alt-CtrlH":      "DeleteWordLeft",
	"Alt-Backspace":  "DeleteWordLeft",
	"Tab":            "Autocomplete|IndentSelection|InsertTab",
	"Backtab":        "CycleAutocompleteBack|OutdentSelection|OutdentLine",
	"Ctrl-o":         "OpenFile",
	"Ctrl-s":         "Save",
	"Ctrl-f":         "Find",
	"Alt-F":          "FindLiteral",
	"Ctrl-n":         "FindNext",
	"Ctrl-p":         "FindPrevious",
	"Alt-[":          "DiffPrevious|CursorStart",
	"Alt-]":          "DiffNext|CursorEnd",
	"Ctrl-z":         "Undo",
	"Ctrl-y":         "Redo",
	"Ctrl-c":         "CopyLine|Copy",
	"Ctrl-x":         "Cut",
	"Ctrl-k":         "CutLine",
	"Ctrl-d":         "DuplicateLine",
	"Ctrl-v":         "Paste",
	"Ctrl-a":         "SelectAll",
	"Ctrl-t":         "AddTab",
	"Alt-,":          "PreviousTab",
	"Alt-.":          "NextTab",
	"Home":           "StartOfTextToggle",
	"End":            "EndOfLine",
	"CtrlHome":       "CursorStart",
	"CtrlEnd":        "CursorEnd",
	"PageUp":         "CursorPageUp",
	"PageDown":       "CursorPageDown",
	"CtrlPageUp":     "PreviousTab",
	"CtrlPageDown":   "NextTab",
	"Ctrl-g":         "ToggleHelp",
	"Alt-g":          "ToggleKeyMenu",
	"Ctrl-r":         "ToggleRuler",
	"Ctrl-l":         "command-edit:goto ",
	"Delete":         "Delete",
	"Ctrl-b":         "ShellMode",
	"Alt-c":          "ShellInsert",
	"Ctrl-q":         "Quit",
	"Ctrl-e":         "CommandMode",
	"Ctrl-w":         "NextSplit",
	"Ctrl-u":         "ToggleMacro",
	"Ctrl-j":         "PlayMacro",
	"Insert":         "ToggleOverwriteMode",

	// Emacs-style keybindings
	"Alt-f": "WordRight",
	"Alt-b": "WordLeft",
	"Alt-a": "StartOfText",
	"Alt-e": "EndOfLine",
	// "Alt-p": "CursorUp",
	// "Alt-n": "CursorDown",

	// Integration with file managers
	"F2":  "Save",
	"F3":  "Find",
	"F4":  "Quit",
	"F7":  "Find",
	"F10": "Quit",
	"Esc": "Escape,Deselect,ClearInfo,RemoveAllMultiCursors,UnhighlightSearch",

	// Mouse bindings
	"MouseWheelUp":     "ScrollUp",
	"MouseWheelDown":   "ScrollDown",
	"MouseLeft":        "MousePress",
	"MouseLeftDrag":    "MouseDrag",
	"MouseLeftRelease": "MouseRelease",
	"MouseMiddle":      "PastePrimary",
	"Ctrl-MouseLeft":   "MouseMultiCursor",

	"Alt-n":        "SpawnMultiCursor",
	"Alt-m":        "SpawnMultiCursorSelect",
	"AltShiftUp":   "SpawnMultiCursorUp",
	"AltShiftDown": "SpawnMultiCursorDown",
	"Alt-p":        "RemoveMultiCursor",
	"Alt-c":        "RemoveAllMultiCursors",
	"Alt-x":        "SkipMultiCursor",
}

var infodefaults = map[string]string{
	"Up":             "HistoryUp",
	"Down":           "HistoryDown",
	"Right":          "CursorRight",
	"Left":           "CursorLeft",
	"ShiftUp":        "SelectUp",
	"ShiftDown":      "SelectDown",
	"ShiftLeft":      "SelectLeft",
	"ShiftRight":     "SelectRight",
	"AltLeft":        "StartOfTextToggle",
	"AltRight":       "EndOfLine",
	"AltUp":          "CursorStart",
	"AltDown":        "CursorEnd",
	"AltShiftRight":  "SelectWordRight",
	"AltShiftLeft":   "SelectWordLeft",
	"CtrlLeft":       "WordLeft",
	"CtrlRight":      "WordRight",
	"CtrlShiftLeft":  "SelectToStartOfTextToggle",
	"ShiftHome":      "SelectToStartOfTextToggle",
	"CtrlShiftRight": "SelectToEndOfLine",
	"ShiftEnd":       "SelectToEndOfLine",
	"CtrlUp":         "CursorStart",
	"CtrlDown":       "CursorEnd",
	"CtrlShiftUp":    "SelectToStart",
	"CtrlShiftDown":  "SelectToEnd",
	"Enter":          "ExecuteCommand",
	"CtrlH":          "Backspace",
	"Backspace":      "Backspace",
	"OldBackspace":   "Backspace",
	"Alt-CtrlH":      "DeleteWordLeft",
	"Alt-Backspace":  "DeleteWordLeft",
	"Tab":            "CommandComplete",
	"Backtab":        "CycleAutocompleteBack",
	"Ctrl-z":         "Undo",
	"Ctrl-y":         "Redo",
	"Ctrl-c":         "CopyLine|Copy",
	"Ctrl-x":         "Cut",
	"Ctrl-k":         "CutLine",
	"Ctrl-v":         "Paste",
	"Home":           "StartOfTextToggle",
	"End":            "EndOfLine",
	"CtrlHome":       "CursorStart",
	"CtrlEnd":        "CursorEnd",
	"Delete":         "Delete",
	"Ctrl-q":         "AbortCommand",
	"Ctrl-e":         "EndOfLine",
	"Ctrl-a":         "StartOfLine",
	"Ctrl-w":         "DeleteWordLeft",
	"Insert":         "ToggleOverwriteMode",
	"Ctrl-b":         "WordLeft",
	"Ctrl-f":         "WordRight",
	"Ctrl-d":         "DeleteWordLeft",
	"Ctrl-m":         "ExecuteCommand",
	"Ctrl-n":         "HistoryDown",
	"Ctrl-p":         "HistoryUp",
	"Ctrl-u":         "SelectToStart",

	// Emacs-style keybindings
	"Alt-f": "WordRight",
	"Alt-b": "WordLeft",
	"Alt-a": "StartOfText",
	"Alt-e": "EndOfLine",

	// Integration with file managers
	"F10": "AbortCommand",
	"Esc": "AbortCommand",

	// Mouse bindings
	"MouseWheelUp":     "HistoryUp",
	"MouseWheelDown":   "HistoryDown",
	"MouseLeft":        "MousePress",
	"MouseLeftDrag":    "MouseDrag",
	"MouseLeftRelease": "MouseRelease",
	"MouseMiddle":      "PastePrimary",
}
