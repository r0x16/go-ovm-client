package ovmHelper

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
