package ovmHelper

// vmDiskMapping element
// The VmDiskMapping object contains the relationships of the VirtualDisk and StorageElement for a Vm.
type Vdm struct {
	VmId                *Id    `json:"vmId,omitempty"`
	VirtualDiskId       *Id    `json:"virtualDiskId,omitempty"`
	DiskTarget          int    `json:"diskTarget"`
	EmulatedBlockDevice bool   `json:"emulatedBlockDevice,omitempty"` // false,
	StorageElementId    *Id    `json:"storageElementId,omitempty"`
	DiskWriteMode       string `json:"diskWriteMode,omitempty"` // "READ_ONLY",READ_WRITE
	Id                  *Id    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`        //: "...",
	Description         string `json:"description,omitempty"` //  : "...",
	Locked              bool   `json:"locked,omitempty"`      // false,
	ReadOnly            bool   `json:"readOnly,omitempty"`    //false,
	Generation          int    `json:"generation,omitempty"`
	ResourceGroupIds    *[]Id  `json:"resourceGroupIds,omitempty"` //" : [ {  }, ... ]
	//	userData            // [ {"key" : "...", "value" : "..." }, ... ],
}

var ListVdm []Vdm
