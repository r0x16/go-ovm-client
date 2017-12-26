package ovmHelper

type Vm struct {
	RepositoryId Id `json:"repositoryId"`
	ServerPoolId Id `json:"serverPoolId"`
	//The following fields of the new Vm are optional:
	//	BootOrder          []string `json:"bootOrder"`
	CpuCount           int    `json:"cpuCount,omitempty"`
	CpuCountLimit      int    `json:"cpuCountLimit,omitempty"`
	CpuPriority        int    `json:"cpuPriority,omitempty"`
	CpuUtilizationCap  int    `json:"cpuUtilizationCap,omitempty"`
	HighAvailability   bool   `json:"highAvailability"`
	HugePagesEnabled   bool   `json:"hugePagesEnabled"`
	KeymapName         string `json:"keymapName,omitempty"`
	Memory             int    `json:"memory,omitempty"`
	MemoryLimit        int    `json:"memoryLimit,omitempty"`
	NetworkInstallPath string `json:"networkInstallPath,omitempty"`
	OsType             string `json:"osType,omitempty"`
	ServerId           string `json:"serverId,omitempty"`
	VmDomainType       string `json:"vmDomainType"`
	VmMouseType        string `json:"vmMouseType,omitempty"`
	VmRunState         string `json:"vmRunState,omitempty"`
	VmStartPolicy      string `json:"vmStartPolicy,omitempty"`
}
