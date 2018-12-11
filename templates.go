package core

import uuid "github.com/satori/go.uuid"

type Request struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        *string                `json:"domain"`
	Locale        *string                `json:"locale"`
	Params        map[string]interface{} `json:"params"`
	ResponseQueue string                 `json:"responseQueue"`
	Source        string                 `json:"source,omitempty"`
	Subscribe     bool                   `json:"subscribe,omitempty"`
	CacheKey      string                 `json:"cacheKey,omitempty"`
	Token         string                 `json:"token,omitempty"`
}

type SuccessResponse struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        string                 `json:"domain"`
	Locale        string                 `json:"locale"`
	Result        map[string]interface{} `json:"result"`
	Source        string                 `json:"source"`
	ResponseQueue string                 `json:"responseQueue"`
	CacheKey      string                 `json:"cacheKey,omitempty"`
	Token         string                 `json:"token,omitempty"`
}

type ErrorResponse struct {
	ID            uuid.UUID              `json:"id"`
	Namespace     string                 `json:"namespace"`
	Method        string                 `json:"method"`
	Domain        string                 `json:"domain"`
	Locale        string                 `json:"locale"`
	Error         map[string]interface{} `json:"error"`
	Source        string                 `json:"source"`
	ResponseQueue string                 `json:"responseQueue"`
	CacheKey      string                 `json:"cacheKey,omitempty"`
	Token         string                 `json:"token,omitempty"`
}
