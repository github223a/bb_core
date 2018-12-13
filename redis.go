package core

type Redis struct {
	Host   string `json:"host"`
	Port   uint32 `json:"port"`
	Prefix string `json:"prefix"`
}
