package ovmHelper

import (
	"fmt"
	"log"
)

type VmcsmService struct {
	client *Client
}

func (v *VmcsmService) Read(vmcsmId string) (*Vmcsm, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneStorageMapping/%s", vmcsmId)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	m := &Vmcsm{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VmcsmService) Create(vmCloneDefinitionId string, vmcsm Vmcsm) (*string, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneStorageMapping", vmCloneDefinitionId)
	req, err := v.client.NewRequest("POST", url, nil, vmcsm)
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

func (v *VmcsmService) Delete(vmCloneDefinitionId string, vmcsmId string) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneStorageMapping/%s", vmCloneDefinitionId, vmcsmId)
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
