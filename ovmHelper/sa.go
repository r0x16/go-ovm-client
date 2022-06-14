package ovmHelper

import (
	"encoding/json"
	"fmt"
	"log"
)

type SaService struct {
	client *Client
}

func (sa *SaService) ReadAllIds() ([]*Id, error) {
	req, err := sa.client.NewRequest("GET", "/ovm/core/wsapi/rest/StorageArray/id", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Id
	_, err = sa.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (sa *SaService) ReadAll() ([]*Sa, error) {
	req, err := sa.client.NewRequest("GET", "/ovm/core/wsapi/rest/StorageArray", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Sa
	_, err = sa.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (sa *SaService) Read(id string) (*Sa, error) {
	req, err := sa.client.NewRequest("GET", "/ovm/core/wsapi/rest/StorageArray/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Sa{}
	_, err = sa.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (sa *SaService) Refresh(saId string) error {
	req, err := sa.client.NewRequest("PUT", "/ovm/core/wsapi/rest/StorageArray/"+saId+"/refresh", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = sa.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	sa.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := sa.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}
