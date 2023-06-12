package service

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/call"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/micro"
)

type Consumption struct {
	WorkspaceUUID string  `json:"workspace_uuid"`
	Credit        float64 `json:"credit"`
}

type SendConsumptionRequest struct {
	Consumption []Consumption `json:"consumption"`
}

func SendConsumptionCron() {
	if len(micro.UsageManager.WorkspaceToConsumption) == 0 {
		return
	}

	var consumptions = make([]Consumption, 0)
	for workspaceUUID, credit := range micro.UsageManager.WorkspaceToConsumption {
		consumptions = append(consumptions, Consumption{
			WorkspaceUUID: workspaceUUID,
			Credit:        credit,
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
	micro.UsageManager.WorkspaceToConsumption = make(map[string]float64)
}
