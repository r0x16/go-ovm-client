package ovmHelper

import (
	"encoding/json"
	"fmt"
	"log"
)

type SeService struct {
	client *Client
}

func (se *SeService) ReadAllIds() ([]*Id, error) {
	req, err := se.client.NewRequest("GET", "/ovm/core/wsapi/rest/StorageElement/id", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Id
	_, err = se.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (se *SeService) ReadAll() ([]*Se, error) {
	req, err := se.client.NewRequest("GET", "/ovm/core/wsapi/rest/StorageElement", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Se
	_, err = se.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (se *SeService) Read(id string) (*Se, error) {
	req, err := se.client.NewRequest("GET", "/ovm/core/wsapi/rest/StorageElement/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Se{}
	_, err = se.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (se *SeService) UpdateSe(seId string, element Se) error {

	rSe, _ := se.Read(seId)
	element.Id = rSe.Id
	element.Generation = rSe.Generation

	req, err := se.client.NewRequest("PUT", "/ovm/core/wsapi/rest/StorageElement/"+seId, nil, element)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = se.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	se.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := se.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}

func (se *SeService) RefreshSe(seId string) error {
	req, err := se.client.NewRequest("PUT", "/ovm/core/wsapi/rest/StorageElement/"+seId+"/refresh", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = se.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	se.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := se.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}
