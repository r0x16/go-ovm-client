package ovmHelper

import (
	"fmt"
)

type VdmService struct {
	client *Client
}

func (v *VdmService) Read(vmId string, vdmId string) (*Vdm, error) {

	listOfVdm, err := v.List(vmId)
	if err != nil {
		return nil, err
	}

	for _, v := range *listOfVdm {
		if vdmId == v.Id.Value {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Did not find vmDiskMapping")

}
func (v *VdmService) List(vmId string) (*[]Vdm, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmDiskMapping", vmId)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &ListVdm
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VdmService) Create(vdm Vdm) (*string, error) {

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmDiskMapping", vdm.VmId.Value)
	req, err := v.client.NewRequest("POST", url, nil, vdm)
	if err != nil {
		fmt.Println("error")
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

	return &j.ResultId.Value, nil
}

func (v *VdmService) Delete(vdm Vdm) error {

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmDiskMapping/%s", vdm.VmId.Value, vdm.Id.Value)
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

	return nil
}
