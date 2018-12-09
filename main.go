package core

import (
	rmq "bb-rmq"
)

var Core = coreStruct{
	Config:         Config{},
	Rabbit:         rmq.RabbitMQ{},
	Infrastructure: Infrastructure{},
	Methods:        methods,
}

type coreStruct struct {
	Config         Config             `json:"config"`
	Rabbit         rmq.RabbitMQ       `json:"rabbit"`
	Infrastructure Infrastructure     `json:"infrastructure"`
	Methods        map[string]*Method `json:"methods"`
}

func (core *coreStruct) Init() {

}
