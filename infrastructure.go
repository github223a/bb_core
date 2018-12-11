package core

type Infrastructure struct {
	RedisPrefix            string                       `json:"redisPrefix"`
	RedisPrefixSession     string                       `json:"redisPrefixSession"`
	RedisPrefixSessionList string                       `json:"redisPrefixSessionList"`
	TokenAlg               string                       `json:"tokenAlg"`
	TokenKey               string                       `json:"tokenKey"`
	SessionLifetime        float64                      `json:"sessionLifetime"`
	Expectation            float64                      `json:"expectation"`
	Sharding               map[string]Sharding          `json:"sharding"`
	Settings               map[string]NamespaceSettings `json:"settings"`
}

type NamespaceSettings struct {
	Methods        map[string]MethodSettings `json:"methods"`
	ResponseQueues []string                  `json:"responseQueues"`
}

type Sharding struct {
}
