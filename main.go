package core

var Data Core

type CoreConfig struct {
	Log           Log      `json:"log"`
	Redis         Redis    `json:"redis"`
	RabbitMQ      RabbitMQ `json:"rabbitMQ"`
}

type Core struct {
	Config CoreConfig `json:"config"`
}

func (*Core) InitConfig() {
	Data.Config = ReadConfig()
}
