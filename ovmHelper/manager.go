package ovmHelper

import "errors"

type ManagerService struct {
	client *Client
}

func (m *ManagerService) Read() (*Manager, error) {
	req, err := m.client.NewRequest("GET", "/ovm/core/wsapi/rest/Manager", nil, nil)
	if err != nil {
		return nil, err
	}

	var list []*Manager
	_, err = m.client.Do(req, &list)

	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errors.New("error: Empty Managers")
	}

	return list[0], nil
}
