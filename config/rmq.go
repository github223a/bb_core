package config

type RabbitMQ struct {
	Connection               Connection                 `json:"connection,omitempty"`
	Channels                 map[string]ChannelSettings `json:"channels"`
	//TestChannelName          string                     `json:"testChannelName,omitempty"`
	//TestChannelPingTimeout   int64                      `json:"testChannelPingTimeout,omitempty"`
	//InfrastructureBindingKey string                     `json:"infrastructureBindingKey,omitempty"`
}

type Connection struct {
	Protocol string `json:"protocol,omitempty" default:"amqp"`
	Hostname string `json:"hostname,omitempty" default:"localhost"`
	Username string `json:"username,omitempty" default:"guest"`
	Password string `json:"password,omitempty" default:"guest"`
	Port     int64  `json:"port,omitempty" default:"5672"`
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
	MessageTTL int64 `json:"messageTtl,omitempty"`
	Durable    bool  `json:"durable,omitempty"`
	NoAck      bool  `json:"noAck,omitempty"`
	AutoDelete bool  `json:"autoDelete,omitempty"`
}

//type RabbitMQ struct {
//	Connection RabbitSettings `json:"connection"`
//}
//
//type RabbitSettings struct {
//	Hostname string `json:"hostname"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Port uint64 `json:"port"`
//}