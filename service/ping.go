package service

import "github.com/metadiv-io/saas/micro"

type Pong struct {
	SystemUUID string `json:"system_uuid"`
	SystemName string `json:"system_name"`
}

func PingHandler() micro.Service[struct{}] {
	return func(ctx *micro.Context[struct{}]) {
		ctx.OK(&Pong{
			SystemUUID: ctx.Engine.GingerEngine.SystemUUID,
			SystemName: ctx.Engine.GingerEngine.SystemName,
		})
	}
}
