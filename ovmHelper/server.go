package ovmHelper

import (
	"encoding/json"
	"fmt"
	"log"
)

type ServerService struct {
	client *Client
}

func (ss *ServerService) ReadAllIds() ([]*Id, error) {
	req, err := ss.client.NewRequest("GET", "/ovm/core/wsapi/rest/Server/id", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Id
	_, err = ss.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (ss *ServerService) ReadAll() ([]*Server, error) {
	req, err := ss.client.NewRequest("GET", "/ovm/core/wsapi/rest/Server", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Server
	_, err = ss.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (ss *ServerService) Read(id string) (*Server, error) {
	req, err := ss.client.NewRequest("GET", "/ovm/core/wsapi/rest/Server/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Server{}
	_, err = ss.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (ss *ServerService) Refresh(serverId string) error {
	req, err := ss.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Server/"+serverId+"/refresh", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = ss.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	ss.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := ss.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}
