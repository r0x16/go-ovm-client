package ovmHelper

import (
	"fmt"
	"log"
)

type VmcdService struct {
	client *Client
}

func (v *VmcdService) Read(vmcdId string) (*Vmcd, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s", vmcdId)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	m := &Vmcd{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VmcdService) Create(VmId string, vmcd Vmcd) (*string, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmCloneDefinition", VmId)
	req, err := v.client.NewRequest("POST", url, nil, vmcd)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] %v", req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)

	return &j.ResultId.Value, err
}

func (v *VmcdService) Delete(vmId string, vmCloneDefinitionId string) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmCloneDefinition/%s", vmId, vmCloneDefinitionId)
	req, err := v.client.NewRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)

	return nil
}
