package ovmHelper

type Sa struct {
	Id                   *Id      `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	Locked               bool     `json:"locked,omitempty"`
	ReadOnly             bool     `json:"readOnly,omitempty"`
	Generation           int      `json:"generation,omitempty"`
	UserData             []string `json:"userData,omitempty"`
	ResourceGroupIds     *[]Id    `json:"resourceGroupIds,omitempty"`
	ZoneIds              *[]Id    `json:"zoneIds,omitempty"`
	AccessGroupIds       *[]Id    `json:"accessGroupIds,omitempty"`
	StorageArrayPluginID *Id      `json:"storageArrayPluginId,omitempty"`
	AdminServerIds       *[]Id    `json:"adminServerIds,omitempty"`
	VolumeGroupIds       *[]Id    `json:"volumeGroupIds,omitempty"`
	ServerID             *Id      `json:"serverId,omitempty"`
	AdminHostname        string   `json:"adminHostname,omitempty"`
	AdminUsername        string   `json:"adminUsername,omitempty"`
	AdminPassword        string   `json:"adminPassword,omitempty"`
	Status               string   `json:"status,omitempty"`
	TotalSize            int      `json:"totalSize,omitempty"`
	UsedSize             int      `json:"usedSize,omitempty"`
	FreeSize             int      `json:"freeSize,omitempty"`
	AllocatedSize        int      `json:"allocatedSize,omitempty"`
	StorageDescription   string   `json:"storageDescription,omitempty"`
	ExtraInformation     string   `json:"extraInformation,omitempty"`
	QosValues            []string `json:"qosValues,omitempty"`
	Validated            bool     `json:"validated,omitempty"`
	DefaultArray         bool     `json:"defaultArray,omitempty"`
	StorageName          string   `json:"storageName,omitempty"`
	KnownStorageNames    []string `json:"knownStorageNames,omitempty"`
	StorageType          string   `json:"storageType,omitempty"`
	AccessHosts          string   `json:"accessHosts,omitempty"`
	UseChap              string   `json:"useChap,omitempty"`
	LipScan              bool     `json:"lipScan,omitempty"`
}
