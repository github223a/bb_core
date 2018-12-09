package core

type Core struct {
	Config         Config            `json:"config"`
	Infrastructure Infrastructure    `json:"infrastructure"`
	Methods        map[string]Method `json:"methods"`
}

func (core *Core) Init() {
}
