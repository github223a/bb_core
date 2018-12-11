package core

import (
	uuid "github.com/satori/go.uuid"
)

type RabbitMQ struct {
	Connection               Connection                 `json:"connection,omitempty"`
	Channels                 map[string]ChannelSettings `json:"channels"`
	TestChannelName          string                     `json:"testChannelName,omitempty"`
	TestChannelPingTimeout   int                        `json:"testChannelPingTimeout,omitempty"`
	InfrastructureBindingKey string                     `json:"infrastructureBindingKey,omitempty"`
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

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"prefix"`
}

type Log struct {
	Type         string `json:"type"`
	FilePath     string `json:"filePath"`
	FileLevel    string `json:"fileLevel"`
	ConsoleLevel string `json:"consoleLevel"`
}

type ChannelSettings struct {
	ExchangeName    string        `json:"exchangeName,omitempty"`
	ExchangeType    string        `json:"exchangeType,omitempty"`
	QueueName       string        `json:"queueName,omitempty"`
	QueueOptions    QueueSettings `json:"queueOptions,omitempty"`
	ConsumeActivate bool          `json:"consumeActivate,omitempty"`
	BindingKey      string        `json:"bindingKey,omitempty"`
}

type QueueSettings struct {
	MessageTTL int32 `json:"messageTtl,omitempty"`
	Durable    bool  `json:"durable,omitempty"`
	NoAck      bool  `json:"noAck,omitempty"`
	AutoDelete bool  `json:"autoDelete,omitempty"`
}

type Request struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        *string                `json:"domain"`
	Locale        *string                `json:"locale"`
	Params        map[string]interface{} `json:"params"`
	ResponseQueue string                 `json:"responseQueue"`
	Source        string                 `json:"source,omitempty"`
	Subscribe     bool                   `json:"subscribe,omitempty"`
	CacheKey      string                 `json:"cacheKey,omitempty"`
	Token         string                 `json:"token,omitempty"`
}

type SuccessResponse struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        string                 `json:"domain"`
	Locale        string                 `json:"locale"`
	Result        map[string]interface{} `json:"result"`
	Source        string                 `json:"source"`
	ResponseQueue string                 `json:"responseQueue"`
	CacheKey      string                 `json:"cacheKey,omitempty"`
	Token         string                 `json:"token,omitempty"`
}

type ErrorResponse struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        string                 `json:"domain"`
	Locale        string                 `json:"locale"`
	Error         map[string]interface{} `json:"error"`
	Source        string                 `json:"source"`
	ResponseQueue string                 `json:"responseQueue"`
	CacheKey      string                 `json:"cacheKey,omitempty"`
	Token         string                 `json:"token,omitempty"`
}
