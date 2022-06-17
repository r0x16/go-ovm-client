package ovmHelper

type Server struct {
	Id                        *Id              `json:"id,omitempty"`
	Name                      string           `json:"name,omitempty"`
	Description               string           `json:"description,omitempty"`
	Locked                    bool             `json:"locked,omitempty"`
	ReadOnly                  bool             `json:"readOnly,omitempty"`
	Generation                int              `json:"generation,omitempty"`
	UserData                  []string         `json:"userData,omitempty"`
	ResourceGroupIds          *[]Id            `json:"resourceGroupIds,omitempty"`
	ServerRunState            string           `json:"serverRunState,omitempty"`
	IPAddress                 string           `json:"ipAddress,omitempty"`
	ServerRoles               []string         `json:"serverRoles,omitempty"`
	MaintenanceMode           bool             `json:"maintenanceMode,omitempty"`
	AgentLogin                string           `json:"agentLogin,omitempty"`
	StatisticInterval         int              `json:"statisticInterval,omitempty"`
	NtpServers                []string         `json:"ntpServers,omitempty"`
	BiosVendor                string           `json:"biosVendor,omitempty"`
	BiosVersion               string           `json:"biosVersion,omitempty"`
	BiosReleaseDate           string           `json:"biosReleaseDate,omitempty"`
	ProcessorType             string           `json:"processorType,omitempty"`
	ProcessorSpeed            float64          `json:"processorSpeed,omitempty"`
	PopulatedProcessorSockets int              `json:"populatedProcessorSockets,omitempty"`
	ThreadsPerCore            int              `json:"threadsPerCore,omitempty"`
	CoresPerProcessorSocket   int              `json:"coresPerProcessorSocket,omitempty"`
	TotalProcessorCores       int              `json:"totalProcessorCores,omitempty"`
	EnabledProcessorCores     int              `json:"enabledProcessorCores,omitempty"`
	Memory                    int              `json:"memory,omitempty"`
	UsableMemory              int              `json:"usableMemory,omitempty"`
	NoExecuteFlag             bool             `json:"noExecuteFlag,omitempty"`
	OvmVersion                string           `json:"ovmVersion,omitempty"`
	Hostname                  string           `json:"hostname,omitempty"`
	ManagerUUID               string           `json:"managerUuid,omitempty"`
	Manufacturer              string           `json:"manufacturer,omitempty"`
	ProductName               string           `json:"productName,omitempty"`
	SerialNumber              string           `json:"serialNumber,omitempty"`
	ControlDomainIds          *[]Id            `json:"controlDomainIds,omitempty"`
	CPUIds                    *[]Id            `json:"cpuIds,omitempty"`
	Hypervisor                *Hypervisor      `json:"hypervisor,omitempty"`
	ServerControllerID        *Id              `json:"serverControllerId,omitempty"`
	CPUCompatibilityGroupID   *Id              `json:"cpuCompatibilityGroupId,omitempty"`
	RepositoryExportIds       *[]Id            `json:"repositoryExportIds,omitempty"`
	VMIds                     *[]Id            `json:"vmIds,omitempty"`
	ServerPoolID              *Id              `json:"serverPoolId,omitempty"`
	EthernetPortIds           *[]Id            `json:"ethernetPortIds,omitempty"`
	ClusterID                 *Id              `json:"clusterId,omitempty"`
	NetworkIds                *[]Id            `json:"networkIds,omitempty"`
	FileSystemMountIds        *[]Id            `json:"fileSystemMountIds,omitempty"`
	FileServerPluginIds       *[]Id            `json:"fileServerPluginIds,omitempty"`
	StorageArrayPluginIds     *[]Id            `json:"storageArrayPluginIds,omitempty"`
	LocalStorageArrayID       *Id              `json:"localStorageArrayId,omitempty"`
	StorageInitiatorIds       *[]Id            `json:"storageInitiatorIds,omitempty"`
	StorageElementIds         *[]Id            `json:"storageElementIds,omitempty"`
	RefreshFileServerIds      *[]Id            `json:"refreshFileServerIds,omitempty"`
	AccessGroupIds            *[]Id            `json:"accessGroupIds,omitempty"`
	LocalFileServerID         *Id              `json:"localFileServerId,omitempty"`
	ServerAbilities           *ServerAbilities `json:"serverAbilities,omitempty"`
	ServerUpToDate            bool             `json:"serverUpToDate,omitempty"`
	RebootOnUpgrade           bool             `json:"rebootOnUpgrade,omitempty"`
	Protected                 bool             `json:"protected,omitempty"`
}

type Hypervisor struct {
	Capabilities []string `json:"capabilities,omitempty"`
	Version      string   `json:"version,omitempty"`
	Type         string   `json:"type,omitempty"`
}

type ServerAbilities struct {
	Cluster                       bool `json:"cluster,omitempty"`
	Nfs                           bool `json:"nfs,omitempty"`
	FibreChannel                  bool `json:"fibreChannel,omitempty"`
	HighAvailability              bool `json:"highAvailability,omitempty"`
	VMSuspend                     bool `json:"vmSuspend,omitempty"`
	PerVMCPUOverSubscribe         bool `json:"perVmCpuOverSubscribe,omitempty"`
	AllVMCPUOverSubscribe         bool `json:"allVmCpuOverSubscribe,omitempty"`
	BondModeActiveBackup          bool `json:"bondModeActiveBackup,omitempty"`
	BondModeLinkAggregation       bool `json:"bondModeLinkAggregation,omitempty"`
	BondModeAdaptiveLoadBalancing bool `json:"bondModeAdaptiveLoadBalancing,omitempty"`
	MtuConfiguration              bool `json:"mtuConfiguration,omitempty"`
	LocalStorageElement           bool `json:"localStorageElement,omitempty"`
	VMMemoryAlignment             int  `json:"vmMemoryAlignment,omitempty"`
	VncConsole                    bool `json:"vncConsole,omitempty"`
	SerialConsole                 bool `json:"serialConsole,omitempty"`
	MigrationSetup                bool `json:"migrationSetup,omitempty"`
	HvmMaxVnics                   int  `json:"hvmMaxVnics,omitempty"`
	ServerPackageUpdate           bool `json:"serverPackageUpdate,omitempty"`
	PowerOnWOL                    bool `json:"powerOnWOL,omitempty"`
	RepositoryOnSharedDisk        bool `json:"repositoryOnSharedDisk,omitempty"`
	RepositoryOnLocalDisk         bool `json:"repositoryOnLocalDisk,omitempty"`
	ClusterFsOnPhysicalDisk       bool `json:"clusterFsOnPhysicalDisk,omitempty"`
	VMEmptyCdrom                  bool `json:"vmEmptyCdrom,omitempty"`
	VMRestartActions              bool `json:"vmRestartActions,omitempty"`
	VMLiveStorageMigration        bool `json:"vmLiveStorageMigration,omitempty"`
	IScsi                         bool `json:"iScsi,omitempty"`
}
