package ovmHelper

import (
	"encoding/json"
	"fmt"
	"log"
)

type FsService struct {
	client *Client
}

func (fss *FsService) ReadAllIds() ([]*Id, error) {
	req, err := fss.client.NewRequest("GET", "/ovm/core/wsapi/rest/FileSystem/id", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Id
	_, err = fss.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (fss *FsService) ReadAll() ([]*Fs, error) {
	req, err := fss.client.NewRequest("GET", "/ovm/core/wsapi/rest/FileSystem", nil, nil)
	if err != nil {
		return nil, err
	}

	var m []*Fs
	_, err = fss.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (fss *FsService) Read(id string) (*Fs, error) {
	req, err := fss.client.NewRequest("GET", "/ovm/core/wsapi/rest/FileSystem/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Fs{}
	_, err = fss.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (fss *FsService) Refresh(fsId string) error {
	req, err := fss.client.NewRequest("PUT", "/ovm/core/wsapi/rest/FileSystem/"+fsId+"/refresh", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = fss.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	fss.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := fss.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}
