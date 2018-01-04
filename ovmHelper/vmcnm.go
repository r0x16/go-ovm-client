package ovmHelper

import (
	"fmt"
	"log"
)

type VmcnmService struct {
	client *Client
}

func (v *VmcnmService) Read(vmcnmId string) (*Vmcnm, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneNetworkMapping/%s", vmcnmId)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	m := &Vmcnm{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VmcnmService) Create(vmCloneDefinitionId string, vmcnm Vmcnm) (*string, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneNetworkMapping", vmCloneDefinitionId)
	req, err := v.client.NewRequest("POST", url, nil, vmcnm)
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

	log.Printf("[DEBUG] %v", j)
	return &j.ResultId.Value, err
}

func (v *VmcnmService) Delete(vmCloneDefinitionId string, vmcnmId string) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneNetworkMapping/%s", vmCloneDefinitionId, vmcnmId)
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
