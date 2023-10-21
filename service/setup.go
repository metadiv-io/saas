package service

import (
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/router"
)

func QuickSetup(engine micro.IEngine) {
	router.GET(engine, "/ping", PingHandler)
	ginger.CRON(engine, "@every 10s", RegisterCron(engine))
	ginger.CRON(engine, "@every 15s", GetMicroIpCron(engine))
	ginger.CRON(engine, "@every 1m", SendConsumptionCron)
}
