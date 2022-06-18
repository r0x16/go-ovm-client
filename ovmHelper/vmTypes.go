package ovmHelper

type Vm struct {
	Id                       *Id      `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Description              string   `json:"description,omitempty"`
	Locked                   bool     `json:"locked,omitempty"`
	ReadOnly                 bool     `json:"readOnly,omitempty"`
	Generation               int      `json:"generation,omitempty"`
	UserData                 []string `json:"userData,omitempty"`
	ResourceGroupIds         *[]Id    `json:"resourceGroupIds,omitempty"`
	VmRunState               string   `json:"vmRunState,omitempty"`
	CpuPriority              int      `json:"cpuPriority,omitempty"`
	Memory                   int      `json:"memory,omitempty"`
	MemoryLimit              int      `json:"memoryLimit,omitempty"`
	HugePagesEnabled         bool     `json:"hugePagesEnabled,omitempty"`
	CpuCount                 int      `json:"cpuCount,omitempty"`
	CpuCountLimit            int      `json:"cpuCountLimit,omitempty"`
	CpuUtilizationCap        int      `json:"cpuUtilizationCap,omitempty"`
	HighAvailability         bool     `json:"highAvailability,omitempty"`
	BootOrder                []string `json:"bootOrder,omitempty"`
	VmMouseType              string   `json:"vmMouseType,omitempty"`
	OsType                   string   `json:"osType,omitempty"`
	VmDomainType             string   `json:"vmDomainType,omitempty"`
	DiskLimit                int      `json:"diskLimit,omitempty"`
	RestartActionOnCrash     string   `json:"restartActionOnCrash,omitempty"`
	RestartActionOnPowerOff  string   `json:"restartActionOnPowerOff,omitempty"`
	RestartActionOnRestart   string   `json:"restartActionOnRestart,omitempty"`
	VmDiskMappingIds         *[]Id    `json:"vmDiskMappingIds,omitempty"`
	VirtualNicIds            *[]Id    `json:"virtualNicIds,omitempty"`
	RepositoryId             *Id      `json:"repositoryId,omitempty"`
	ServerPoolId             *Id      `json:"serverPoolId,omitempty"`
	NetworkInstallPath       string   `json:"networkInstallPath,omitempty"`
	VmCloneDefinitionIds     *[]Id    `json:"vmCloneDefinitionIds,omitempty"`
	Origin                   string   `json:"origin,omitempty"`
	KeymapName               string   `json:"keymapName,omitempty"`
	Viridian                 bool     `json:"viridian,omitempty"`
	ServerId                 *Id      `json:"serverId,omitempty"`
	SslVncPort               string   `json:"sslVncPort,omitempty"`
	SslTtyPort               string   `json:"sslTtyPort,omitempty"`
	AffinityGroupIds         *[]Id    `json:"affinityGroupIds,omitempty"`
	OsVersion                string   `json:"osVersion,omitempty"`
	KernelVersion            string   `json:"kernelVersion,omitempty"`
	GuestDriverVersion       string   `json:"guestDriverVersion,omitempty"`
	VmApiVersion             string   `json:"vmApiVersion,omitempty"`
	VmStartPolicy            string   `json:"vmStartPolicy,omitempty"`
	CurrentMemory            int      `json:"currentMemory,omitempty"`
	CurrentCpuCount          int      `json:"currentCpuCount,omitempty"`
	CurrentDomainId          int      `json:"currentDomainId,omitempty"`
	VmConfigFileAbsolutePath string   `json:"vmConfigFileAbsolutePath,omitempty"`
	VmConfigFileMountedPath  string   `json:"vmConfigFileMountedPath,omitempty"`
	Architecture             string   `json:"architecture,omitempty"`
}

type CfgVm struct {
	NetworkId    string
	SendMessages *[]KeyValuePair
	RootPassword *[]KeyValuePair
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
