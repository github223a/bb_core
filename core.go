package core

var Data CoreStructure
var Infrastructure InfrastructureStructure

func (*CoreStructure) InitCore() {
	initConfig()
	initInfrastructure()
}

func initConfig() {
	Data.Config = ReadConfig()
}

func initInfrastructure() {
	Data.Infrastructure = InfrastructureStructure{
		RedisPrefixSession:     "sess:",
		RedisPrefixSessionList: "sesslist:",
		TokenAlg:               "HS256",
		SessionLifetime:        86400,
		Expectation:            10000,
	}
}

type CoreStructure struct {
	Config         interface{}             `json:"config"`
	Infrastructure InfrastructureStructure `json:"infrastructure"`
}

type CommonConfig struct {
	Log         Log      `json:"log"`
	RabbitMQ    RabbitMQ `json:"rabbitMQ"`
	PingTimeout int32    `json:"pingTimeout"`
}
