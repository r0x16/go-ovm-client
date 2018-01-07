package ovmHelper

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type VmService struct {
	client *Client
}

func (v *VmService) Read(id string) (*Vm, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vm{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VmService) Stop(id string) error {
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+id+"/stop", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return err
	}

	for v.client.Jobs.Running(m.Id.Value) {
		fmt.Println("true")
		time.Sleep(5 * time.Second)
	}

	return err
}

func (v *VmService) Start(id string) error {
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+id+"/start", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return err
	}

	for v.client.Jobs.Running(m.Id.Value) {
		time.Sleep(5 * time.Second)
	}
	return err
}

func (v *VmService) CreateVm(vm Vm, cfgVm CfgVm) (*string, error) {
	req, err := v.client.NewRequest("POST", "/ovm/core/wsapi/rest/Vm", nil, vm)
	if err != nil {
		s := ""
		return &s, err
	}

	log.Printf("[DEBUG] vmrequest: %v", req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)

	if err != nil {
		s := ""
		return &s, err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)

	j, _ := v.client.Jobs.Read(m.Id.Value)
	if cfgVm.NetworkId != "" {
		vn := Vn{NetworkId: &Id{Type: "com.oracle.ovm.mgr.ws.model.Network",
			Value: cfgVm.NetworkId},
			VmId: &Id{Type: "com.oracle.ovm.mgr.ws.model.Vm",
				Value: j.ResultId.Value,
			},
		}
		_, err = v.client.Vns.Create(j.ResultId.Value, vn)
		if err != nil {
			return &j.ResultId.Value, err
		}
	}

	return &j.ResultId.Value, err
}

func (v *VmService) CloneVm(cloneVmId string, vmCloneDefinitionId string, vm Vm) (*string, error) {

	params := make(map[string]string)
	params["vmId"] = cloneVmId
	params["serverPoolId"] = vm.ServerPoolId.Value
	params["repositoryId"] = vm.RepositoryId.Value
	if vmCloneDefinitionId != "" {
		params["vmCloneDefinitionId"] = vmCloneDefinitionId
	}

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/clone", cloneVmId)
	req, err := v.client.NewRequest("PUT", url, params, nil)
	if err != nil {
		return nil, err
	}

	m := &JobResponse{}

	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if err != nil {
		return nil, err
	}

	jsonJobResult, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jsonJobResult))
	vm.Id = &Id{Value: j.ResultId.Value,
		Type: "com.oracle.ovm.mgr.ws.model.Vm"}

	log.Printf("[INFO] Update vmId: %s", j.ResultId.Value)
	err = v.client.Vms.UpdateVm(j.ResultId.Value, vm)

	return &j.ResultId.Value, err
}

func (v *VmService) UpdateVm(vmId string, vm Vm) error {
	p := make(map[string]string)

	p["vmId"] = vmId
	rVm, _ := v.client.Vms.Read(vmId)
	if rVm.CpuCount != vm.CpuCount {
		rVm.CpuCount = vm.CpuCount
	}
	if rVm.CpuCountLimit != vm.CpuCountLimit {
		rVm.CpuCountLimit = vm.CpuCountLimit
	}
	if rVm.Name != vm.Name {
		rVm.Name = vm.Name
	}
	if rVm.VmDomainType != vm.VmDomainType {
		rVm.VmDomainType = vm.VmDomainType
	}
	if rVm.Memory != vm.Memory {
		rVm.Memory = vm.Memory
	}
	if rVm.HugePagesEnabled != vm.HugePagesEnabled {
		rVm.HugePagesEnabled = vm.HugePagesEnabled
	}
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+vmId, p, rVm)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}

func (v *VmService) DeleteVm(vmId string) error {

	req, err := v.client.NewRequest("DELETE", "/ovm/core/wsapi/rest/Vm/"+vmId, nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	//j, _ := v.client.Jobs.Read(m.Id.Value)
	return err
}
