package ovmHelper

type Vn struct {
	MacAddress string `json:"macAddress,omitempty"`
	//IpAddresses   *[]Id  `json:"ipAddresses,omitempty"`
	//InterfaceName string `json:"interfaceName,omitempty"`
	VmId      *Id    `json:"vmId,omitempty"`
	NetworkId *Id    `json:"networkId,omitempty"`
	Id        *Id    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	//Description      string `json:description,omitempty"`
	//Locked           bool   `json:locked,omitempty"`
	//ReadOnly         bool   `json:readOnly,omitempty"`
	Generation int `json:"generation,omitempty"`
	//ResourceGroupIds *[]Id  `json:resourceGroupIds,omitempty"`
}
