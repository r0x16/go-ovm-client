package ovmHelper

import (
	"fmt"
	"time"
)

type VmService struct {
	client *Client
}

func (v *VmService) Read(id string) (*VmResponse, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &VmResponse{}
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

func (v *VmService) CreateVm(vm Vm) error {
	req, err := v.client.NewRequest("POST", "/ovm/core/wsapi/rest/Vm", nil, vm)
	if err != nil {
		return err
	}

	fmt.Println(req)

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
