package ovmHelper

type Vd struct {
	DiskType              string `json:"diskType,omitempty"` // "VIRTUAL_DISK",
	Size                  int    `json:"size,omitempty"`
	OnDiskSize            int    `json:"onDiskSize,omitempty"`
	Path                  string `json:"path,omitempty"`
	VmDiskMappingIds      []*Id  `json:"vmDiskMappingIds,omitempty"`
	RepositoryId          *Id    `json:"repositoryId,omitempty"`
	Shareable             bool   `json:"shareable,omitempty"`
	ImportFileName        string `json:"importFileName,omitempty"`
	AbsolutePath          string `json:"absolutePath,omitempty"`
	MountedPath           string `json:"mountedPath,omitempty"`
	AssemblyVirtualDiskId *Id    `json:"assemblyVirtualDiskId,omitempty"`
	Id                    *Id    `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	Description           string `json:"description,omitempty"`
	Locked                bool   `json:"locked,omitempty"`   //false,
	ReadOnly              bool   `json:"readOnly,omitempty"` //: false,
	Generation            int    `json:"generation,omitempty"`
	/*"userData" : [ {
	  "key" : "...",
	  "value" : "..."
	}, ... ],*/
	ResourceGroupIds []*Id `json:"resourceGroupIds,omitempty"`
}
