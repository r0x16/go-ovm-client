package ovmHelper

type id struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Uri   string `json:"uri"`
	Name  string `json:"name"`
}

type VmResponse struct {
	//RepositoryId int64 `json:"repositoryId"`
	//serverPoolId int64
	Id   id     `json:"id"`
	Name string `json:"name"`
	//The following fields of the new Vm are optional:
	//BootOrder          []string
	CpuCount                 int    `json:"cpuCount"`
	CpuCountLimit            int    `json:"cpuCountLimit"`
	CpuPriority              int    `json:"cpuPriority"`
	CpuUtilizationCap        int    `json:"cpuUtilizationCap"`
	HighAvailability         bool   `json:"highAvailability"`
	HugePagesEnabled         bool   `json:"hugePagesEnabled"`
	KeymapName               string `json:"keymapName"`
	Memory                   int    `json:"memory"`
	MemoryLimit              int    `json:"memoryLimit"`
	NetworkInstallPath       string `json:"networkInstallPath"`
	OsType                   string `json:"osType"`
	VmDomainType             string `json:"vmDomainType"`
	VmMouseType              string `json:"vmMouseType"`
	VmRunState               string `json:"vmRunState"`
	VmStartPolicy            string `json:"vmStartPolicy"`
	SslVncPort               int    `json:""sslVncPort"`
	SslTtyPort               string `json:""sslTtyPort"`
	AffinityGroupIds         []int  `json:""affinityGroupIds"`
	OsVersion                string `json:""osVersion"`
	KernelVersion            string `json:""kernelVersion"`
	GuestDriverVersion       string `json:""guestDriverVersion"`
	VmApiVersion             string `json:""vmApiVersion"`
	CurrentMemory            int    `json:""currentMemory"`
	CurrentCpuCount          int    `json:""currentCpuCount"`
	CurrentDomainId          int    `json:""currentDomainId"`
	VmConfigFileAbsolutePath string `json:""vmConfigFileAbsolutePath"`
	VmConfigFileMountedPath  string `json:""vmConfigFileMountedPath"`
	Architecture             string `json:""architecture"`
	VmDiskMappingIds         []id   `json:"vmDiskMappingIds"`
	VirtualNicIds            []id   `json:"virtualNicIds"`
	RepositoryId             id     `json:"repositoryId"`
	ServerPoolId             id     `json:"serverPoolId"`
	ServerId                 id     `json:"serverId"`
}

type JobResponse struct {
	Id   id   `json:"id"`
	Done bool `json:"done"`
	/*	ResulitId                    id     `json:"resourceGroupIds"`

		SummaryDone                  bool   `json:"summaryDone"`
		JobGroup                     bool   `json:"jobGroup"`
		JobRunState                  string `json:"jobRunState"`
		JobSummaryState              string `json:"jobSummaryState"`
		AbortedByUser                string `json:"abortedByUser"`
		ExtraInfo                    string `json:"extraInfo"`
		Name                         string `json:"name"`
		Description                  string `json:"description"`
		Locked                       bool   `json:"locked"`
		ReadOnly                     bool   `json:"readOnly"`
		Generation                   int    `json:"generation"`
		ProgressMessage              string `json:"progressMessage"`
		LatestSummaryProgressMessage string `json:"latestSummaryProgressMessage"`
		StartTime                    int64  `json:"startTime"`
		EndTime                      int64  `json:"endTime"`
		ParentJobId                  id     `json:"parentJobId"`
		ChildJobIds                  id     `json:"childJobIds"`
		User                         string `json:"user"`
		ResourceGroupIds             id     `json:"resourceGroupIds"`*/
}
