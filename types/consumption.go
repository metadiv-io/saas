package types

type Consumption struct {
	WorkspaceUUID string  `json:"workspace_uuid"`
	UserUUID      string  `json:"user_uuid"`
	Credit        float64 `json:"credit"`
}
