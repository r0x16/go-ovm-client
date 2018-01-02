package ovmHelper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	User     string
	Password string
	BaseURL  *url.URL
	client   *http.Client
	Vms      *VmService
	Vds      *VdService
	Vdms     *VdmService
	Jobs     *JobService
}

func NewClient(user string, password string, baseUrl string) *Client {
	baseURL, _ := url.Parse(baseUrl)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &Client{User: user, Password: password, BaseURL: baseURL}
	c.client = &http.Client{Transport: tr}
	c.Vms = &VmService{client: c}
	c.Vds = &VdService{client: c}
	c.Vdms = &VdmService{client: c}
	c.Jobs = &JobService{client: c}
	return c
}

func (pc *Client) NewRequest(method string, rsc string, params map[string]string, data interface{}) (*http.Request, error) {
	fmt.Println("inside NewRewuest func")
	baseUrl, err := url.Parse(pc.BaseURL.String() + rsc)
	if err != nil {
		return nil, err
	}

	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			fmt.Printf("value: %s \n", v)
			ps.Set(k, v)
		}
		baseUrl.RawQuery = ps.Encode()
	}
	req, err := http.NewRequest(method, baseUrl.String(), nil)
	if data != nil {
		jsonString, _ := json.Marshal(data)
		req, err = http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(jsonString))
	}
	req.SetBasicAuth(pc.User, pc.Password)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func (pc *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := pc.client.Do(req)
	if err != nil {
		fmt.Println("Do request failed")
		return nil, err
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("status: %s \n", resp.Status)
	//fmt.Println(string(body))

	/*if err := validateResponse(resp); err != nil {
		return resp, err
	}


	*/
	err = decodeResponse(resp, v)
	return resp, err

}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	err := json.Unmarshal([]byte(bodyString), &v)
	//err := json.Unmarshal([]byte(bodyBytes), &v)
	return err
}
