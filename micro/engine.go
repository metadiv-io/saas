package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Engine struct {
	Gin        *gin.Engine
	Cron       *cron.Cron
	SystemUUID string
	SystemName string
	PubPEM     string // Public key for JWT
}

func NewEngine(systemUUID, systemName string) *Engine {
	return &Engine{
		Gin:        gin.New(),
		Cron:       cron.New(),
		SystemUUID: systemUUID,
		SystemName: systemName,
	}
}

func (e *Engine) CRON(spec string, job func()) {
	e.Cron.AddFunc(spec, job)
}

func (e *Engine) Run(addr ...string) error {
	go e.Cron.Start()
	return e.Gin.Run(addr...)
}
