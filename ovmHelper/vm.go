package ovmHelper

import (
	"fmt"
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
		fmt.Println("true")
		time.Sleep(5 * time.Second)
	}
	return err
}

func (v *VmService) CreateVm(vm Vm) (*string, error) {
	req, err := v.client.NewRequest("POST", "/ovm/core/wsapi/rest/Vm", nil, vm)
	if err != nil {
		s := ""
		return &s, err
	}

	fmt.Println(req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)

	if err != nil {
		s := ""
		return &s, err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)

	j, _ := v.client.Jobs.Read(m.Id.Value)
	return &j.ResultId.Value, err
}

func (v *VmService) UpdateVm(vmId string, vm Vm) (*string, error) {
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
		s := ""
		return &s, err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		s := ""
		fmt.Println("inside error")
		return &s, err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	return &j.ResultId.Value, err
}

func (v *VmService) DeleteVm(vmId string) (*string, error) {

	req, err := v.client.NewRequest("DELETE", "/ovm/core/wsapi/rest/Vm/"+vmId, nil, nil)
	if err != nil {
		s := ""
		return &s, err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		s := ""
		fmt.Println("inside error")
		return &s, err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	return &j.ResultId.Value, err
}
