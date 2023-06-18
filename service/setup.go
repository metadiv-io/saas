package service

import (
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/router"
)

func QuickSetup(engine *micro.Engine) {
	router.GET(engine, "/ping", PingHandler)
	engine.CRON("@every 10s", RegisterCron(engine))
	engine.CRON("@every 15s", GetMicroIpCron(engine))
	engine.CRON("@every 1m", SendConsumptionCron)
}
