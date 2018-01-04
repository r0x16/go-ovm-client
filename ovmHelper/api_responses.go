package ovmHelper

import "fmt"

type Id struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Uri   string `json:"uri,omitempty"`
	Name  string `json:"name,omitempty"`
}

type VmResponse struct {
	//RepositoryId int64 `json:"repositoryId"`
	//serverPoolId int64
	Id   Id     `json:"id"`
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
	SslVncPort               int    `json:"sslVncPort"`
	SslTtyPort               string `json:"sslTtyPort"`
	AffinityGroupIds         []int  `json:"affinityGroupIds"`
	OsVersion                string `json:"osVersion"`
	KernelVersion            string `json:"kernelVersion"`
	GuestDriverVersion       string `json:"guestDriverVersion"`
	VmApiVersion             string `json:"vmApiVersion"`
	CurrentMemory            int    `json:"currentMemory"`
	CurrentCpuCount          int    `json:"currentCpuCount"`
	CurrentDomainId          int    `json:"currentDomainId"`
	VmConfigFileAbsolutePath string `json:"vmConfigFileAbsolutePath"`
	VmConfigFileMountedPath  string `json:"vmConfigFileMountedPath"`
	Architecture             string `json:"architecture"`
	VmDiskMappingIds         []Id   `json:"vmDiskMappingIds"`
	VirtualNicIds            []Id   `json:"virtualNicIds"`
	RepositoryId             Id     `json:"repositoryId"`
	ServerPoolId             Id     `json:"serverPoolId"`
	ServerId                 Id     `json:"serverId"`
}

type JobResponse struct {
	Id                           *Id    `json:"id"`
	Done                         bool   `json:"done,omitempty"`
	ResultId                     *Id    `json:"resultId,omitempty"`
	ResourceGroupIds             *[]Id  `json:"resourceGroupIds,omitempty"`
	SummaryDone                  bool   `json:"summaryDone,omitempty"`
	JobGroup                     bool   `json:"jobGroup,omitempty"`
	JobRunState                  string `json:"jobRunState,omitempty"`
	JobSummaryState              string `json:"jobSummaryState,omitempty"`
	AbortedByUser                string `json:"abortedByUser,omitempty"`
	ExtraInfo                    string `json:"extraInfo,omitempty"`
	Name                         string `json:"name,omitempty"`
	Description                  string `json:"description,omitempty"`
	Locked                       bool   `json:"locked,omitempty"`
	ReadOnly                     bool   `json:"readOnly,omitempty"`
	Generation                   int    `json:"generation,omitempty"`
	ProgressMessage              string `json:"progressMessage,omitempty"`
	LatestSummaryProgressMessage string `json:"latestSummaryProgressMessage,omitempty"`
	StartTime                    int64  `json:"startTime,omitempty"`
	EndTime                      int64  `json:"endTime,omitempty"`
	ParentJobId                  *Id    `json:"parentJobId,omitempty"`
	ChildJobIds                  *[]Id  `json:"childJobIds,omitempty"`
	User                         string `json:"user,omitempty"`
}

type OvmHelperError struct {
	StatusCode int    `json:"statuscode"`
	StatusDesc string `json:"statusdesc"`
	Message    string `json:"errormessage"`
}

func (r *OvmHelperError) Error() string {
	return fmt.Sprintf("%d %v: %v", r.StatusCode, r.StatusDesc, r.Message)
}

type errorJsonResponse struct {
	Error *OvmHelperError `json:"error"`
}
