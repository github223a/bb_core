package core

type RabbitMQ struct {
	Connection               Connection                 `json:"connection,omitempty"`
	Channels                 map[string]ChannelSettings `json:"channels"`
	TestChannelName          string                     `json:"testChannelName,omitempty"`
	TestChannelPingTimeout   int                        `json:"testChannelPingTimeout,omitempty"`
	InfrastructureBindingKey string                     `json:"infrastructureBindingKey,omitempty"`
}

type Connection struct {
	Protocol string `json:"protocol,omitempty" default:"amqp"`
	Hostname string `json:"hostname,omitempty" default:"localhost"`
	Username string `json:"username,omitempty" default:"guest"`
	Password string `json:"password,omitempty" default:"guest"`
	Port     int    `json:"port,omitempty" default:"5672"`
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
	MessageTTL int32 `json:"messageTTL,omitempty"`
	Durable    bool  `json:"durable,omitempty"`
	NoAck      bool  `json:"noack,omitempty"`
	AutoDelete bool  `json:"autoDelete,omitempty"`
}
