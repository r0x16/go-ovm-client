package ovmHelper

type Fs struct {
	Id                       Id       `json:"id"`
	Name                     string   `json:"name"`
	Description              string   `json:"description"`
	Locked                   bool     `json:"locked"`
	ReadOnly                 bool     `json:"readOnly"`
	Generation               int      `json:"generation"`
	UserData                 []string `json:"userData"`
	ResourceGroupIds         []Id     `json:"resourceGroupIds"`
	FileServerIds            []Id     `json:"fileServerIds"`
	FileSystemMountIds       []Id     `json:"fileSystemMountIds"`
	RepositoryIds            []Id     `json:"repositoryIds"`
	StorageElementID         Id       `json:"storageElementId"`
	ClusterID                Id       `json:"clusterId"`
	ClusterHeartbeatDeviceID Id       `json:"clusterHeartbeatDeviceId"`
	Path                     string   `json:"path"`
	Size                     int64    `json:"size"`
	FreeSize                 int64    `json:"freeSize"`
	Refreshed                bool     `json:"refreshed"`
	AccessGroupID            Id       `json:"accessGroupId"`
	Shared                   bool     `json:"shared"`
}
