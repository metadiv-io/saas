package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/saas/constant"
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
	constant.CheckEnvSet(SYSTEM_UUID, "SYSTEM_UUID is required")

	SYSTEM_NAME := env.String("SYSTEM_NAME")
	constant.CheckEnvSet(SYSTEM_NAME, "SYSTEM_NAME is required")

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
