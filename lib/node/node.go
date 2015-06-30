package node

import (
	"crypto/rsa"
	"fmt"
	client "github.com/paulmooring/chef-swarm/lib/client"
	"io"
	"io/ioutil"
	"net/http"
)

func place_holder() {
	fmt.Println("hi")
}

type Node struct {
	Name      string
	ServerURL string
	Key       *rsa.PrivateKey
	Client    *http.Client
}

func NewNode(name string, server_url string, key_file string) (Node, error) {
	key, err := client.LoadKey(key_file)
	if err != nil {
		return Node{}, err
	}

	node := Node{
		Name:      name,
		ServerURL: server_url,
		Key:       key,
		Client:    &http.Client{},
	}

	return node, nil
}

func (nd Node) request(method string, path string, body io.Reader) (int, []byte, error) {
	http_req, err := http.NewRequest(method, nd.ServerURL+path, body)
	err = client.PrepRequest(http_req, nd.Name, nd.Key)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	resp, err := nd.Client.Do(http_req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return status, resp_body, nil
}

func (nd Node) Get(path string) (int, []byte, error) {
	status, resp_body, err := nd.request("GET", path, nil)
	if err != nil {
		return 0, nil, err
	}

	return status, resp_body, nil
}

func (nd Node) Post(path string, body io.Reader) (int, []byte, error) {
	status, resp_body, err := nd.request("POST", path, body)
	if err != nil {
		return 0, nil, err
	}

	return status, resp_body, nil
}

func (nd Node) Put(path string, body io.Reader) (int, []byte, error) {
	status, resp_body, err := nd.request("PUT", path, body)
	if err != nil {
		return 0, nil, err
	}

	return status, resp_body, nil
}
