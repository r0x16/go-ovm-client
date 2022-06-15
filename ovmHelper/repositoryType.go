package ovmHelper

type Repository struct {
	Id                 Id       `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
	Locked             bool     `json:"locked,omitempty"`
	ReadOnly           bool     `json:"readOnly,omitempty"`
	Generation         int      `json:"generation,omitempty"`
	UserData           []string `json:"userData,omitempty"`
	ResourceGroupIds   []Id     `json:"resourceGroupIds,omitempty"`
	VirtualDiskIds     []Id     `json:"virtualDiskIds,omitempty"`
	VMIds              []Id     `json:"vmIds,omitempty"`
	AssemblyIds        []Id     `json:"assemblyIds,omitempty"`
	PresentedServerIds []Id     `json:"presentedServerIds,omitempty"`
	FileSystemID       Id       `json:"fileSystemId,omitempty"`
	ManagerUUID        string   `json:"managerUuid,omitempty"`
	Refreshed          bool     `json:"refreshed,omitempty"`
	SharePath          string   `json:"sharePath,omitempty"`
}
