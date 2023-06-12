package micro

import (
	"time"

	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/saas/call"
	"github.com/metadiv-io/saas/constant"
	"github.com/metadiv-io/saas/types"
)

var UsageManager = &usageManager{
	WorkspaceToAllowed:     make(map[string]map[string]bool),
	WorkspaceToExpired:     make(map[string]map[string]time.Time),
	WorkspaceToConsumption: make(map[string]float64),
	UUIDToApi:              make(map[string]*types.Api),
	TagToApi:               make(map[string]*types.Api),
}

type usageManager struct {
	WorkspaceToAllowed     map[string]map[string]bool      // workspace_uuid -> api_uuid -> allowed
	WorkspaceToExpired     map[string]map[string]time.Time // workspace_uuid -> api_uuid -> expired
	WorkspaceToConsumption map[string]float64              // workspace_uuid -> consumption
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
	ApiUUID       string `json:"api_uuid"`
}

type IsAllowResponse struct {
	Allowed bool `json:"allowed"`
}

func (m *usageManager) AskWorkspaceAllowed(workspaceUUID string, apiUUID string) bool {
	w, ok1 := m.WorkspaceToAllowed[workspaceUUID]
	if ok1 {
		_, ok2 := w[apiUUID]
		if ok1 && ok2 {
			if time.Now().Before(m.WorkspaceToExpired[workspaceUUID][apiUUID]) {
				return m.WorkspaceToAllowed[workspaceUUID][apiUUID]
			}
		}
	}

	resp, err := call.POST[IsAllowResponse](nil, constant.MICRO_SERVICE_HOST_AUTH, "/micro/allowed", IsAllowRequest{
		WorkspaceUUID: workspaceUUID,
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
		if m.WorkspaceToAllowed[workspaceUUID] == nil {
			m.WorkspaceToAllowed[workspaceUUID] = make(map[string]bool)
			m.WorkspaceToExpired[workspaceUUID] = make(map[string]time.Time)
		}
		m.WorkspaceToAllowed[workspaceUUID][apiUUID] = true
		m.WorkspaceToExpired[workspaceUUID][apiUUID] = time.Now().Add(time.Hour)
		m.WorkspaceToConsumption[workspaceUUID] += m.UUIDToApi[apiUUID].Credit
		return true
	}

	return false
}
