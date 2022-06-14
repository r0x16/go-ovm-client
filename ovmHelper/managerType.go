package ovmHelper

type Manager struct {
	Id               *Id      `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Description      string   `json:"description,omitempty"`
	Locked           bool     `json:"locked,omitempty"`
	ReadOnly         bool     `json:"readOnly,omitempty"`
	Generation       int      `json:"generation,omitempty"`
	UserData         []string `json:"userData,omitempty"`
	ResourceGroupIds *[]Id    `json:"resourceGroupIds,omitempty"`
	ManagerUUID      string   `json:"managerUuid,omitempty"`
	ManagerVersion   string   `json:"managerVersion,omitempty"`
	ManagerRunState  string   `json:"managerRunState,omitempty"`
	ManagerTime      int64    `json:"managerTime,omitempty"`
}
