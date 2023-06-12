package service

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/call"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/types"
)

type RegisterRequest struct {
	SystemUUID string                `json:"system_uuid"`
	SystemName string                `json:"system_name"`
	UsageApi   map[string]*types.Api `json:"usage_api"`
}

type RegisterResponse struct {
	PublicPEM string                `json:"public_pem"`
	UsageApi  map[string]*types.Api `json:"usage_api"`
}

func RegisterCron(engine *micro.Engine) func() {
	return func() {
		resp, err := call.POST[RegisterResponse](nil, constant.MICRO_SERVICE_HOST_AUTH, "/micro/register", &RegisterRequest{
			SystemUUID: engine.SystemUUID,
			SystemName: engine.SystemName,
			UsageApi:   micro.UsageManager.UUIDToApi,
		}, nil)
		if err != nil {
			logger.Error("Fail to register service:", err)
		}
		if !resp.Success {
			logger.Error("Fail to register service:", resp.Error.Message)
		}
		engine.PubPEM = resp.Data.PublicPEM
		micro.UsageManager.UUIDToApi = resp.Data.UsageApi
		for _, api := range micro.UsageManager.UUIDToApi {
			micro.UsageManager.TagToApi[api.Tag()] = api
		}
	}
}
