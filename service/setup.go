package service

import (
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/router"
)

func QuickSetup(engine *micro.Engine) {
	router.GET(engine, "/ping", PingHandler)
	ginger.CRON(engine.GingerEngine, "@every 10s", RegisterCron(engine))
	ginger.CRON(engine.GingerEngine, "@every 15s", GetMicroIpCron(engine))
	ginger.CRON(engine.GingerEngine, "@every 1m", SendConsumptionCron)
}
