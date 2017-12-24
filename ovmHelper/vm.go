package ovmHelper

type Vm struct {
	client *Client
}

func (v *Vm) Read(id string) (*VmResponse, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/"+id, nil)
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
