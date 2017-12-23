package ovmHelper

import "strconv"

type Vm struct {
	client *Client
}

func (v *Vm) Read(id int) (*CheckResponse, error) {
	req, err := cs.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	m := &vmDetailsJsonResponse{}
	_, err = vm.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m.Check, err
}
