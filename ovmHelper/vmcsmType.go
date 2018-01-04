package ovmHelper

type Vmcsm struct {
	VmDiskMappingId     *Id    `json:"vmDiskMappingId,omitempty"`
	VmCloneDefinitionId *Id    `json:"vmCloneDefinitionId,omitempty"`
	RepositoryId        *Id    `json:"repositoryId,omitempty"`
	StorageArrayId      *Id    `json:"storageArrayId,omitempty"`
	StorageElementId    *Id    `json:",omitempty"`
	Id                  *Id    `json:"id,omitempty"`
	CloneType           string `json:"cloneType,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Locked              bool   `json:"locked,imitempty"`
	ReadOnly            bool   `json:"readOnly,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	ResourceGroupIds    *[]Id  `json:"resourceGroupIds,omitempty"`
}
