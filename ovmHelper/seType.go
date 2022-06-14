package ovmHelper

type Se struct {
	Id                       *Id      `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	Locked                   bool     `json:"locked,omitempty"`
	ReadOnly                 bool     `json:"readOnly,omitempty"`
	Generation               int      `json:"generation,omitempty"`
	ResourceGroupIds         *[]Id    `json:"resourceGroupIds,omitempty"`
	AccessGroupIds           *[]Id    `json:"accessGroupIds,omitempty"`
	FileSystemIds            *[]Id    `json:"fileSystemIds,omitempty"`
	StoragePathIds           *[]Id    `json:"storagePathIds,omitempty"`
	ReservingServerIds       *[]Id    `json:"reservingServerIds,omitempty"`
	VMDiskMappingIds         *[]Id    `json:"vmDiskMappingIds,omitempty"`
	VolumeGroupID            *Id      `json:"volumeGroupId,omitempty"`
	StorageArrayID           *Id      `json:"storageArrayId,omitempty"`
	ClusterHeartbeatDeviceID *Id      `json:"clusterHeartbeatDeviceId,omitempty"`
	DeviceNames              []string `json:"deviceNames,omitempty"`
	ExtraInformation         string   `json:"extraInformation,omitempty"`
	Page83ID                 string   `json:"page83Id,omitempty"`
	Qos                      string   `json:"qos,omitempty"`
	ServerReserved           bool     `json:"serverReserved,omitempty"`
	Shareable                bool     `json:"shareable,omitempty"`
	Size                     int64    `json:"size,omitempty"`
	State                    string   `json:"state,omitempty"`
	Status                   string   `json:"status,omitempty"`
	StorageTargetNames       []string `json:"storageTargetNames,omitempty"`
	ThinProvision            bool     `json:"thinProvision,omitempty"`
	Type                     string   `json:"type,omitempty"`
	UserFriendlyName         string   `json:"userFriendlyName,omitempty"`
	Vendor                   string   `json:"vendor,omitempty"`
}
