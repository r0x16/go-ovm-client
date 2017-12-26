package ovmHelper

type Vm struct {
	RepositoryId string `json:"repositoryId"`
	ServerPoolId string `json:"serverPoolId"`
	//The following fields of the new Vm are optional:
	BootOrder          []string
	CpuCount           int    `json:"cpuCount"`
	CpuCountLimit      int    `json:"cpuCountLimit"`
	CpuPriority        int    `json:"cpuPriority"`
	CpuUtilizationCap  int    `json:"cpuUtilizationCap"`
	HighAvailability   bool   `json:"highAvailability"`
	HugePagesEnabled   bool   `json:"hugePagesEnabled"`
	KeymapName         string `json:"keymapName"`
	Memory             int    `json:"memory"`
	MemoryLimit        int    `json:"memoryLimit"`
	NetworkInstallPath string `json:"networkInstallPath"`
	OsType             string `json:"osType"`
	ServerId           string `json:"serverId"`
	VmDomainType       string `json:"vmDomainType"`
	VmMouseType        string `json:"vmMouseType"`
	VmRunState         string `json:"vmRunState"`
	VmStartPolicy      string `json:"vmStartPolicy"`
}
