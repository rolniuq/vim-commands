package edit

const (
	EditName = "edit"
)

type Edit struct{}

func NewEdit() *Edit {
	return &Edit{}
}

func (s *Edit) GetCommands() map[string]string {
	return map[string]string{
		"r":        "replace a single character",
		"R":        "replace more than one character, util ESC pressed",
		"J":        "join line below to current one with one space in between",
		"gJ":       "join line below to current one without space in between",
		"gwip":     "reflow paragraph",
		"g~":       "switch case up to motion",
		"gu":       "change to lowercase up to motion",
		"gU":       "change to uppercase up to motion",
		"cc":       "change entire line",
		"c$ or C":  "replace to the end of the line",
		"ciw":      "change entire word",
		"cw or ce": "change to the end of word",
		"s":        "delete character and substitute text (same as cl)",
		"S":        "delete line and substitute text (same as cc)",
		"xp":       "transpose two letters (delete and paste)",
		"u":        "undo",
		"U":        "restore last changed line",
		"Ctrl+u":   "redo",
		".":        "repeat last command",
	}
}
