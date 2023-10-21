package micro

import (
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/ginger"
)

type IEngine interface {
	ginger.IEngine
	PubPEM() string
	SetPubPEM(pubPEM string)
	MicroIPs() []string
	SetMicroIPs(microIPs []string)
}

type Engine struct {
	ginger.IEngine
	pubPEM   string   // Public key for JWT
	microIPs []string // IP addresses of micro services
}

func (e *Engine) PubPEM() string {
	return e.pubPEM
}

func (e *Engine) SetPubPEM(pubPEM string) {
	e.pubPEM = pubPEM
}

func (e *Engine) MicroIPs() []string {
	return e.microIPs
}

func (e *Engine) SetMicroIPs(microIPs []string) {
	e.microIPs = microIPs
}

func NewEngine() IEngine {
	return &Engine{
		IEngine:  ginger.NewEngine(env.MustString("SYSTEM_UUID"), env.MustString("SYSTEM_NAME")),
		pubPEM:   "",
		microIPs: make([]string, 0),
	}
}
