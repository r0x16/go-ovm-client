package ovmHelper

import (
	"fmt"
	"log"
)

type VnService struct {
	client *Client
}

func (v *VnService) Read(id string) (*Vn, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/VirtualNic/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vn{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VnService) Create(vmId string, vn Vn) (*string, error) {
	params := make(map[string]string)
	params["vmid"] = vmId
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VirtualNic", vmId)
	req, err := v.client.NewRequest("POST", url, params, vn)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] req: %v \n", req)

	m := &JobResponse{}

	_, err = v.client.Do(req, m)
	if err != nil {
		log.Printf("[ERROR] err: %v", err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)

	return &j.ResultId.Value, err
}

func (v *VnService) Delete(vnId string, vn Vn) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VirtualNic/%s", vn.VmId.Value, vnId)
	req, err := v.client.NewRequest("DELETE", url, nil, nil)
	if err != nil {
		fmt.Println("error")
		return err
	}

	m := &JobResponse{}

	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)

	return err
}
