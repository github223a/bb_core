package core

// import (
// 	rmq "bb_rmq"
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	uuid "github.com/satori/go.uuid"
// )

// var Methods = map[string]*Method{
// 	"friendship":     friendship,
// 	"infrastructure": infrastructure,
// }

// var friendship = createMethod(runFriendship, friendShipMethodSettings)
// var infrastructure = createMethod(runInfrastructure, infrastructureMethodSettings)

// func createMethod(run func(transport rmq.RabbitMQ, request rmq.Request), settings MethodSettings) *Method {
// 	return &Method{
// 		Run:      run,
// 		Settings: settings,
// 	}
// }

// func runFriendship(transport rmq.RabbitMQ, request rmq.Request) {
// 	if request.Namespace == NAMESPACE_INTERNAL {
// 		return
// 	}
// 	// handshakeMsg := rmq.Request{
// 	// 	ID:        generateId(),
// 	// 	Namespace: NAMESPACE_INTERNAL,
// 	// 	Method:    HANDSHAKE,
// 	// 	Params:    handshakeParams,
// 	// 	Source:    Core.Config.Namespace,
// 	// }

// 	// transport.SendToInternal(handshakeMsg)
// }

// func runInfrastructure(transport rmq.RabbitMQ, request rmq.Request) {
// 	var info Infrastructure

// 	vbyte, _ := json.Marshal(request.Params)
// 	json.Unmarshal(vbyte, &info)

// 	Data.Infrastructure = info

// 	log.Printf("Infrastructure updated.")
// }

// var infrastructureMethodSettings = MethodSettings{
// 	IsInternal: true,
// 	Auth:       false,
// 	Cache:      0,
// 	Middlewares: Middlewares{
// 		Before: []string{},
// 		After:  []string{},
// 	},
// }

// var friendShipMethodSettings = MethodSettings{
// 	IsInternal: true,
// 	Auth:       false,
// 	Cache:      0,
// 	Middlewares: Middlewares{
// 		Before: []string{},
// 		After:  []string{},
// 	},
// }

// var handshakeParams = map[string]interface{}{
// 	"namespace": Data.Config.Namespace,
// 	"methods": map[string]interface{}{
// 		"friendship":     friendShipMethodSettings,
// 		"infrastructure": infrastructureMethodSettings,
// 	},
// }

// func generateId() uuid.UUID {
// 	id, err := uuid.NewV4()

// 	if err != nil {
// 		fmt.Printf("Something went wrong with generate id: %s", err)
// 		return uuid.UUID{}
// 	}
// 	return id
// }

type Method struct {
	Run      func(transport RabbitMQ, request Request) `json:"run"`
	Settings MethodSettings                            `json:"settings"`
}

type MethodSettings struct {
	IsInternal  bool        `json:"isInternal"`
	Auth        bool        `json:"auth"`
	Cache       float64     `json:"cache"`
	Middlewares Middlewares `json:"middlewares"`
}

type Middlewares struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}
