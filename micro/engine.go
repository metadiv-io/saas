package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/robfig/cron"
)

type Engine struct {
	Gin        *gin.Engine
	Cron       *cron.Cron
	SystemUUID string
	SystemName string
	PubPEM     string   // Public key for JWT
	MicroIPs   []string // IP addresses of micro services
}

func NewEngine() *Engine {
	SYSTEM_UUID := env.String("SYSTEM_UUID")
	if SYSTEM_UUID == "" {
		panic("SYSTEM_UUID is empty")
	}

	SYSTEM_NAME := env.String("SYSTEM_NAME")
	if SYSTEM_NAME == "" {
		panic("SYSTEM_NAME is empty")
	}
	return &Engine{
		Gin:        gin.New(),
		Cron:       cron.New(),
		SystemUUID: SYSTEM_UUID,
		SystemName: SYSTEM_NAME,
	}
}

func (e *Engine) CRON(spec string, job func()) {
	e.Cron.AddFunc(spec, job)
}

func (e *Engine) Run(addr ...string) error {
	go e.Cron.Start()
	return e.Gin.Run(addr...)
}
