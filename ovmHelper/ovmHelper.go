package ovmHelper

import (
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
	Vms      *Vm
	Jobs     *Job
}

func NewClient(user string, password string, baseUrl string) *Client {
	baseURL, _ := url.Parse(baseUrl)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &Client{User: user, Password: password, BaseURL: baseURL}
	//c.client = http.DefaultClient{Transport: tr}
	c.client = &http.Client{Transport: tr}
	c.Vms = &Vm{client: c}
	c.Jobs = &Job{client: c}
	return c
}

func (pc *Client) NewRequest(method string, rsc string, params map[string]string) (*http.Request, error) {
	baseUrl, err := url.Parse(pc.BaseURL.String() + rsc)

	if err != nil {
		return nil, err
	}

	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			ps.Set(k, v)
		}
		baseUrl.RawQuery = ps.Encode()
	}

	req, err := http.NewRequest(method, baseUrl.String(), nil)
	req.SetBasicAuth(pc.User, pc.Password)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	return req, err
}

func (pc *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//	body, err := ioutil.ReadAll(resp.Body)

	//	fmt.Println(string(body))

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
	return err
}
