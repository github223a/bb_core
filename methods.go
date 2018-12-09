package core

import (
	rmq "bb-rmq"
	"encoding/json"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

var methods = map[string]*Method{
	"friendship":     friendship,
	"infrastructure": infrastructure,
}

var friendship = createMethod(runFriendship, friendShipMethodSettings)
var infrastructure = createMethod(runInfrastructure, infrastructureMethodSettings)

func createMethod(run func(request rmq.Request), settings MethodSettings) *Method {
	return &Method{
		Run:      run,
		Settings: settings,
	}
}

func runFriendship(request rmq.Request) {
	if request.Namespace == NAMESPACE_INTERNAL {
		return
	}
	handshakeMsg := rmq.Request{
		ID:        generateId(),
		Namespace: NAMESPACE_INTERNAL,
		Method:    HANDSHAKE,
		Params:    handshakeParams,
		Source:    Core.Config.Namespace,
	}

	// Core.Rabbit.(handshakeMsg)
}

func runInfrastructure(request rmq.Request) {
	var infrastr map[string]NamespaceSettings

	vbyte, _ := json.Marshal(request.Params["infrastructure"])
	json.Unmarshal(vbyte, &infrastr)

	Core.Infrastructure = Infrastructure{
		RedisPrefix:            request.Params["redisPrefix"].(string),
		RedisPrefixSession:     request.Params["redisPrefixSession"].(string),
		RedisPrefixSessionList: request.Params["redisPrefixSessionList"].(string),
		TokenAlg:               request.Params["tokenAlg"].(string),
		TokenKey:               request.Params["tokenKey"].(string),
		SessionLifetime:        request.Params["sessionLifetime"].(float64),
		Expectation:            request.Params["expectation"].(float64),
		Sharding:               request.Params["shardings"].(map[string]Sharding),
		Settings:               infrastr,
	}
	log.Printf("%sInfrastructure updated.")
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

var friendShipMethodSettings = MethodSettings{
	IsInternal: true,
	Auth:       false,
	Cache:      0,
	Middlewares: Middlewares{
		Before: []string{},
		After:  []string{},
	},
}

var handshakeParams = map[string]interface{}{
	"namespace": Core.Config.Namespace,
	"methods": map[string]interface{}{
		"friendship":     friendShipMethodSettings,
		"infrastructure": infrastructureMethodSettings,
	},
}

func generateId() uuid.UUID {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong with generate id: %s", err)
		return uuid.UUID{}
	}
	return id
}

type Method struct {
	Run      func(request rmq.Request) `json:"run"`
	Settings MethodSettings            `json:"settings"`
}

type MethodSettings struct {
	IsInternal  bool        `json:"isInternal"`
	Auth        bool        `json:"auth"`
	Cache       int         `json:"cache"`
	Middlewares Middlewares `json:"middlewares"`
}

type Middlewares struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}
