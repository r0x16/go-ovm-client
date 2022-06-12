package ovmHelper

type Vmcd struct {
	VmId                     *Id    `json:"vmId,omitempty"`
	VmCloneNetworkMappingIds *[]Id  `json:"vmCloneNetworkMappingIds,omitempty"`
	VmCloneStorageMappingIds *[]Id  `json:"vmCloneStorageMappingIds,omitempty"`
	Id                       *Id    `json:"id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	Locked                   bool   `json:"locked,omitempty"`
	ReadOnly                 bool   `json:"readOnly,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	ResourceGroupIds         *[]Id  `json:"resourceGroupIds,omitempty"`
	/*"userData" : [ {
	    "key" : "...",
	    "value" : "..."
	  }, ... ],
	*/

}
