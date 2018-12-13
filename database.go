package core

type Database struct {
	Connections map[string]DatabaseSettings `json:"connections,omitempty"`
	PingTimeout int64                       `json:"pingTimeout,omitempty"`
}

type DatabaseSettings struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}
