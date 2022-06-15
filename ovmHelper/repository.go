package ovmHelper

import (
	"encoding/json"
	"fmt"
	"log"
)

type RepositoryService struct {
	client *Client
}

func (rs *RepositoryService) ReadAllIds() ([]*Id, error) {
	req, err := rs.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository/id", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Id
	_, err = rs.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (rs *RepositoryService) ReadAll() ([]*Repository, error) {
	req, err := rs.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Repository
	_, err = rs.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (rs *RepositoryService) Read(id string) (*Repository, error) {
	req, err := rs.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Repository{}
	_, err = rs.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (rs *RepositoryService) Refresh(repositoryId string) error {
	req, err := rs.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Repository/"+repositoryId+"/refresh", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = rs.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	rs.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := rs.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}
