package ovmHelper

type Fs struct {
	Id                       Id       `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	Locked                   bool     `json:"locked,omitempty"`
	ReadOnly                 bool     `json:"readOnly,omitempty"`
	Generation               int      `json:"generation,omitempty"`
	UserData                 []string `json:"userData,omitempty"`
	ResourceGroupIds         []Id     `json:"resourceGroupIds,omitempty"`
	FileServerIds            []Id     `json:"fileServerIds,omitempty"`
	FileSystemMountIds       []Id     `json:"fileSystemMountIds,omitempty"`
	RepositoryIds            []Id     `json:"repositoryIds,omitempty"`
	StorageElementID         Id       `json:"storageElementId,omitempty"`
	ClusterID                Id       `json:"clusterId,omitempty"`
	ClusterHeartbeatDeviceID Id       `json:"clusterHeartbeatDeviceId,omitempty"`
	Path                     string   `json:"path,omitempty"`
	Size                     int64    `json:"size,omitempty"`
	FreeSize                 int64    `json:"freeSize,omitempty"`
	Refreshed                bool     `json:"refreshed,omitempty"`
	AccessGroupID            Id       `json:"accessGroupId,omitempty"`
	Shared                   bool     `json:"shared,omitempty"`
}
