package core

type InfrastructureStructure struct {
	RedisPrefix            string                        `json:"redisPrefix"`
	RedisPrefixSession     string                        `json:"redisPrefixSession"`
	RedisPrefixSessionList string                        `json:"redisPrefixSessionList"`
	TokenAlg               string                        `json:"tokenAlg"`
	TokenKey               string                        `json:"tokenKey"`
	SessionLifetime        int64                         `json:"sessionLifetime"`
	Expectation            int64                         `json:"expectation"`
	Sharding               map[string]Sharding           `json:"sharding"`
	Data                   map[string]InfrastructureData `json:"data"`
}

type InfrastructureData struct {
	Methods        map[string]MethodSettings `json:"methods"`
	ResponseQueues []string                  `json:"responseQueues"`
}

type Sharding struct {
}
