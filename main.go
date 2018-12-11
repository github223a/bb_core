package core

var Data Core

type Core struct {
	Config interface{} `json:"config"`
}

// func (c *Core) InitCore() {
// 	c.initConfig()
// }

func (*Core) InitConfig() {
	Data.Config = ReadConfig()
}
