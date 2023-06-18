package micro

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/call"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/types"
)

var UsageManager = &usageManager{
	WorkspaceToConsumption: make(map[string]*types.Consumption),
	UUIDToApi:              make(map[string]*types.Api),
	TagToApi:               make(map[string]*types.Api),
}

type usageManager struct {
	WorkspaceToConsumption map[string]*types.Consumption // subscription_uuid -> consumption
	UUIDToApi              map[string]*types.Api
	TagToApi               map[string]*types.Api
}

func (m *usageManager) Register(method, path, uuid string) {
	api := types.Api{
		Method: method,
		Path:   path,
		UUID:   uuid,
	}
	m.UUIDToApi[uuid] = &api
	m.TagToApi[api.Tag()] = &api
}

func (m *usageManager) UpdateCredit(uuid string, credit float64) {
	m.UUIDToApi[uuid].Credit = credit
	m.TagToApi[m.UUIDToApi[uuid].Tag()].Credit = credit
}

func (m *usageManager) GetByUUID(uuid string) *types.Api {
	return m.UUIDToApi[uuid]
}

func (m *usageManager) GetByTag(tag string) *types.Api {
	return m.TagToApi[tag]
}

type IsAllowRequest struct {
	WorkspaceUUID string `json:"workspace_uuid"`
	UserUUID      string `json:"user_uuid"`
	ApiUUID       string `json:"api_uuid"`
}

type IsAllowResponse struct {
	WorkspaceUUID    string  `json:"workspace_uuid"`
	UserUUID         string  `json:"user_uuid"`
	SubscriptionUUID string  `json:"subscription_uuid"`
	Credit           float64 `json:"credit"`
	Allowed          bool    `json:"allowed"`
}

func (m *usageManager) AskWorkspaceAllowed(workspaceUUID, userUUID string, apiUUID string) bool {
	resp, err := call.POST[IsAllowResponse](nil, constant.MICRO_SERVICE_HOST_AUTH, "/micro/allowed", IsAllowRequest{
		WorkspaceUUID: workspaceUUID,
		UserUUID:      userUUID,
		ApiUUID:       apiUUID,
	}, nil)
	if err != nil {
		logger.Error("asking workspace allowed: ", err)
		return false
	}

	if !resp.Success {
		logger.Error("asking workspace allowed: ", resp.Error.Message)
		return false
	}

	if resp.Data.Allowed {
		if resp.Data.Credit > 0 {
			_, ok := m.WorkspaceToConsumption[resp.Data.SubscriptionUUID]
			if !ok {
				m.WorkspaceToConsumption[userUUID] = &types.Consumption{
					SubscriptionUUID: resp.Data.SubscriptionUUID,
					UserUUID:         resp.Data.UserUUID,
					WorkspaceUUID:    resp.Data.WorkspaceUUID,
					Credit:           resp.Data.Credit,
				}
			} else {
				m.WorkspaceToConsumption[userUUID].Credit += resp.Data.Credit
			}
		}
		return true
	}

	return false
}
