package ovmHelper

type Vn struct {
	Id               *Id      `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Description      string   `json:"description,omitempty"`
	Locked           bool     `json:"locked,omitempty"`
	ReadOnly         bool     `json:"readOnly,omitempty"`
	Generation       int      `json:"generation,omitempty"`
	UserData         *[]Id    `json:"userData,omitempty"`
	ResourceGroupIds *[]Id    `json:"resourceGroupIds,omitempty"`
	MacAddress       string   `json:"macAddress,omitempty"`
	IPAddresses      []string `json:"ipAddresses,omitempty"`
	InterfaceName    string   `json:"interfaceName,omitempty"`
	VmId             *Id      `json:"vmId,omitempty"`
	NetworkId        *Id      `json:"networkId,omitempty"`
}
