package core

import (
	rmq "bb_rmq"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Namespace     string    `json:"namespace,omitempty"`
	UseCache      bool      `json:"useCache"`
	UseIsInternal bool      `json:"useIsInternal"`
	Database      *Database `json:"database,omitempty"`
	RabbitMQ      RabbitMQ  `json:"rabbitMQ,omitempty"`
	Redis         *Redis    `json:"redis,omitempty"`
	Location      *Location `json:"location,omitempty"`
}

func (c *Config) Init(url string) {
	configFile, err := os.Open(url)
	FailOnError(err, "Error on open config file")
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)
	uErr := json.Unmarshal([]byte(bytes), &c)

	data, err := json.MarshalIndent(c, ",", " ")
	fmt.Printf("%+v\n", string(data))

	FailOnError(uErr, "Error on unmarhal config file.")
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s", msg)
		panic(err)
	}
}

type RabbitMQ struct {
	Connection               Connection                     `json:"connection,omitempty"`
	Channels                 map[string]rmq.ChannelSettings `json:"channels,omitempty"`
	TestChannelName          string                         `json:"testChannelName,omitempty"`
	TestChannelPingTimeout   int                            `json:"testChannelPingTimeout,omitempty"`
	InfrastructureBindingKey string                         `json:"infrastructureBindingKey,omitempty"`
}

type Database struct {
	Connections map[string]DatabaseSettings `json:"connections,omitempty"`
	PingTimeout uint32                      `json:"pingTimeout,omitempty"`
}

type DatabaseSettings struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type Connection struct {
	Protocol string `json:"protocol,omitempty" default:"amqp"`
	Hostname string `json:"hostname,omitempty" default:"localhost"`
	Username string `json:"username,omitempty" default:"guest"`
	Password string `json:"password,omitempty" default:"guest"`
	Port     int    `json:"port,omitempty" default:"5672"`
}

type Ws struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Path string `json:"path"`
}

type Rest struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Path string `json:"path"`
}

type Location struct {
	Ws   Ws   `json:"ws"`
	Rest Rest `json:"rest"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"prefix"`
}
