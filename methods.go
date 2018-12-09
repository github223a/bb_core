package core

import rmq "bb-rmq"

var List = map[string]*Method{
	"friendship":     friendship,
	"infrastructure": infrastructure,
}

var friendship = createMethod(runFriendship, friendShipMethodSettings)

func runFriendship(request.Request) {
	if request.Namespace == NAMESPACE_INTERNAL {
		return
	}
	HandshakeMsg.Id = generateId()
	sendToInternal(HandshakeMsg)
}

var friendShipMethodSettings = structures.MethodSettings{
	IsInternal: true,
	Auth:       false,
	Cache:      0,
	Middlewares: structures.Middlewares{
		Before: []string{},
		After:  []string{},
	},
}

type Method struct {
	Run      func(request rmq.Request) `json:"run"`
	Settings MethodSettings            `json:"settings"`
}

type MethodSettings struct {
	IsInternal   bool        `json:"isInternal"`
	Auth         bool        `json:"auth"`
	Cache        int         `json:"cache"`
	ForSubscribe bool        `json:"forSubscribe"`
	Middlewares  Middlewares `json:"middlewares"`
}

type Middlewares struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

func createMethod(run func(request structures.Request), settings structures.MethodSettings) *structures.Method {
	return &structures.Method{
		Run:      run,
		Settings: settings,
	}
}
