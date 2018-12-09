package methods

import (
	rmq "bb-rmq"
	"encoding/json"
	"log"
)

var infrastructure = createMethod(runInfrastructure, infrastructureMethodSettings)

func runInfrastructure(request rmq.Request) {
	var infrastr map[string]NamespaceSettings

	vbyte, _ := json.Marshal(request.Params["infrastructure"])
	json.Unmarshal(vbyte, &infrastr)

	InfrastructureData = InfrastructureData{
		RedisPrefix:            request.Params["redisPrefix"].(string),
		RedisPrefixSession:     request.Params["redisPrefixSession"].(string),
		RedisPrefixSessionList: request.Params["redisPrefixSessionList"].(string),
		TokenAlg:               request.Params["tokenAlg"].(string),
		TokenKey:               request.Params["tokenKey"].(string),
		SessionLifetime:        request.Params["sessionLifetime"].(float64),
		Expectation:            request.Params["expectation"].(float64),
		Shardings:              request.Params["shardings"].(map[string]interface{}),
		Infrastructure:         infrastr,
	}
	log.Printf("%sInfrastructure updated.", HEADER_RMQ_MESSAGE)
}

var infrastructureMethodSettings = MethodSettings{
	IsInternal: true,
	Auth:       false,
	Cache:      0,
	Middlewares: Middlewares{
		Before: []string{},
		After:  []string{},
	},
}
