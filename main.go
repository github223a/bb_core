package core

import (
	rmq "bb_rmq"
)

var Data Core

type Core struct {
	Config         Config             `json:"config"`
	Rabbit         rmq.RabbitMQ       `json:"rabbit"`
	Infrastructure Infrastructure     `json:"infrastructure"`
	Methods        map[string]*Method `json:"methods"`
}
