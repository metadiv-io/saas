package service

import "github.com/metadiv-io/saas/micro"

type Pong struct {
	SystemUUID string `json:"system_uuid"`
	SystemName string `json:"system_name"`
}

func PingHandler() micro.HandlerResponse[struct{}] {
	return micro.HandlerResponse[struct{}]{
		Service: func(ctx *micro.Context[struct{}]) {
			ctx.OK(&Pong{
				SystemUUID: ctx.Engine.SystemUUID,
				SystemName: ctx.Engine.SystemName,
			})
		},
	}
}
