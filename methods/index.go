package methods

import rmq "bb-rmq"

var List = map[string]*Method{
	"friendship":     friendship,
	"infrastructure": infrastructure,
}

func createMethod(run func(request rmq.Request), settings MethodSettings) *Method {
	return &Method{
		Run:      run,
		Settings: settings,
	}
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
