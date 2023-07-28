package micro

import (
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/ginger"
)

type Engine struct {
	GingerEngine *ginger.Engine
	PubPEM       string   // Public key for JWT
	MicroIPs     []string // IP addresses of micro services
}

func NewEngine() *Engine {
	return &Engine{
		GingerEngine: ginger.NewEngine(env.MustString("SYSTEM_UUID"), env.MustString("SYSTEM_NAME")),
		PubPEM:       "",
		MicroIPs:     make([]string, 0),
	}
}
