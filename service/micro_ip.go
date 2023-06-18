package service

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/call"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/micro"
)

type MicroIpResponse struct {
	Ips []string `json:"ips"`
}

func GetMicroIpCron(engine *micro.Engine) func() {
	return func() {
		resp, err := call.GET[MicroIpResponse](nil, constant.MICRO_SERVICE_HOST_AUTH, "/micro/ips", nil, nil)
		if err != nil {
			logger.Error("get micro ips failed: ", err)
			return
		}
		if !resp.Success {
			logger.Error("get micro ips failed: ", resp.Error.Message)
			return
		}
		engine.MicroIPs = resp.Data.Ips
	}
}
