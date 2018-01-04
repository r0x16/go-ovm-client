package ovmHelper

type Vmcnm struct {
	NetworkId           *Id    `json:"networkId,omitempty"`
	VmCloneDefinitionId *Id    `json:"vmCloneDefinitionId,omitempty"`
	VirtualNicId        *Id    `json:"virtualNicId,omitempty"`
	Id                  *Id    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Locked              bool   `json:"locked,imitempty"`
	ReadOnly            bool   `json:"readOnly,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	ResourceGroupIds    *[]Id  `json:"resourceGroupIds,omitempty"`
}
