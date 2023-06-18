package service

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/call"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/micro"
	"github.com/metadiv-io/saas/types"
)

type SendConsumptionRequest struct {
	Consumption []types.Consumption `json:"consumption"`
}

func SendConsumptionCron() {
	if len(micro.UsageManager.WorkspaceToConsumption) == 0 {
		return
	}

	var consumptions = make([]types.Consumption, 0)
	for _, consumption := range micro.UsageManager.WorkspaceToConsumption {
		consumptions = append(consumptions, types.Consumption{
			SubscriptionUUID: consumption.SubscriptionUUID,
			WorkspaceUUID:    consumption.WorkspaceUUID,
			UserUUID:         consumption.UserUUID,
			Credit:           consumption.Credit,
		})
	}

	resp, err := call.POST[any](nil, constant.MICRO_SERVICE_HOST_AUTH, "/micro/consumption", &SendConsumptionRequest{
		Consumption: consumptions,
	}, nil)
	if err != nil {
		logger.Error("send consumption failed: ", err)
		return
	}
	if !resp.Success {
		logger.Error("send consumption failed: ", resp.Error.Message)
		return
	}

	logger.Info("sent consumption records: ", len(micro.UsageManager.WorkspaceToConsumption))
	micro.UsageManager.WorkspaceToConsumption = make(map[string]*types.Consumption)
}
