package service

import (
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/types"
)

type Pong struct {
	SystemUUID string      `json:"system_uuid"`
	SystemName string      `json:"system_name"`
	Apis       []types.Api `json:"apis"`
}

func PingHandler() micro.Service[struct{}] {
	return func(ctx micro.IContext[struct{}]) {
		apis := make([]types.Api, 0)
		for _, api := range micro.UsageManager.UUIDToApi {
			apis = append(apis, *api)
		}
		ctx.OK(&Pong{
			SystemUUID: ctx.Engine().SystemUUID(),
			SystemName: ctx.Engine().SystemName(),
			Apis:       apis,
		})
	}
}
