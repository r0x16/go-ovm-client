package ovmHelper

import (
	"fmt"
	"log"
	"strconv"
)

type VdService struct {
	client *Client
}

func (v *VdService) Read(id string) (*Vd, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/VirtualDisk/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vd{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VdService) Create(repositoryId string, sparse bool, vd Vd) (*string, error) {
	params := make(map[string]string)
	params["repositoryId"] = repositoryId
	params["sparse"] = strconv.FormatBool(sparse)
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Repository/%s/VirtualDisk", repositoryId)
	req, err := v.client.NewRequest("POST", url, params, vd)
	if err != nil {
		fmt.Println("error")
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
	if !j.succeed() {
		return nil, j.Error
	}

	return &j.ResultId.Value, err
}

func (v *VdService) Update(vdId string, vd Vd) error {
	params := make(map[string]string)
	params["virtualDiskId"] = vdId

	rVd, _ := v.client.Vds.Read(vdId)
	if rVd.Name != vd.Name {
		rVd.Name = vd.Name
	}
	if rVd.Description != vd.Description {
		rVd.Description = vd.Description
	}
	if rVd.Shareable != vd.Shareable {
		rVd.Shareable = vd.Shareable
	}

	url := fmt.Sprintf("/ovm/core/wsapi/rest/VirtualDisk/%s", vdId)
	req, err := v.client.NewRequest("PUT", url, params, rVd)
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
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}

func (v *VdService) Delete(repositoryId string, vdId string) error {
	params := make(map[string]string)
	params["repositoryId"] = repositoryId
	params["virtualDiskId"] = vdId
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Repository/%s/VirtualDisk/%s", repositoryId, vdId)
	req, err := v.client.NewRequest("DELETE", url, params, nil)
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
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}
