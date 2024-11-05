package edit

type Edit struct{}

func NewEdit() *Edit {
	return &Edit{}
}

func (s *Edit) GetCommands() map[string]string {
	return map[string]string{
		"r": "replace a single character",
		"R": "replace more than one character, util ESC pressed",
		"J": "join line below to current one with one space in between",
	}
}
