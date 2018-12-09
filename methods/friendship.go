package methods

import (
	rmq "bb-rmq"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

var friendship = createMethod(runFriendship, friendShipMethodSettings)

func runFriendship(request rmq.Request) {
	if request.Namespace == NAMESPACE_INTERNAL {
		return
	}
	HandshakeMsg.ID = generateId()
	sendToInternal(HandshakeMsg)
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
	"namespace": core.Namespace,
	"methods": map[string]interface{}{
		"friendship":     friendShipMethodSettings,
		"infrastructure": infrastructureMethodSettings,
	},
}

var HandshakeMsg = Request{
	Namespace: "internal",
	Method:    "handshake",
	Domain:    nil,
	Locale:    nil,
	Params:    handshakeParams,
	Source:    core.Namespace,
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func generateId() uuid.UUID {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong with generate id: %s", err)
		return uuid.UUID{}
	}
	return id
}
