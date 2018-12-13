package core

type Redis struct {
	Host   string `json:"host"`
	Port   int64 `json:"port"`
	Prefix string `json:"prefix"`
}
