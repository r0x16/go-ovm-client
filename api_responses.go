package ovmHelpr

type vmDetailsJsonResponse struct {
	repositoryId int64
	serverPoolId int64
	//The following fields of the new Vm are optional:
	bootOrder          []string
	cpuCount           int
	cpuCountLimit      int
	cpuPriority        int
	cpuUtilizationCap  int
	highAvailability   bool
	hugePagesEnabled   bool
	keymapName         sting
	memory             int
	memoryLimit        int
	networkInstallPath string
	osType             string
	serverId           int64
	vmDomainType       string
	vmMouseType        string
	vmRunState         string
	vmStartPolicy      string
}
