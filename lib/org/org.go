package org

import (
	"encoding/json"
	"fmt"
	config "github.com/paulmooring/chef-swarm/lib/config"
	node "github.com/paulmooring/chef-swarm/lib/node"
	"io"
	"os"
	"sync"
)

type Org struct {
	ServerURI string
	Config    config.OrgConfig
	Validator node.Node
}

func NewOrg(server_uri string, cfg config.OrgConfig) (Org, error) {
	validator, err := node.NewNode(cfg.ValidatorName, server_uri, cfg.ValidatorKey)
	if err != nil {
		return Org{}, err
	}

	o := Org{
		server_uri,
		cfg,
		validator,
	}

	return o, nil
}

func (o Org) CreateNode(name string) error {
	body := node.ClientBody(name)
	_, resp_body, err := o.Validator.Post("/clients", body)
	if err != nil {
		return err
	}

	var dat map[string]interface{}
	err = json.Unmarshal(resp_body, &dat)
	if err != nil {
		return err
	}

	priv_key := dat["private_key"].(string)
	filename := o.Config.KeyDirectory + "/" + name + ".pem"

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.WriteString(file, priv_key)
	if err != nil {
		return err
	}
	file.Close()

	node_client, err := node.NewNode(name, o.ServerURI, filename)
	if err != nil {
		return err
	}

	body = node.NodeBody(name)
	_, _, err = node_client.Post("/nodes", body)
	if err != nil {
		return err
	}

	return nil
}

func (o Org) ChefRuns() error {
	nodes := make([]string, o.Config.Nodes)

	for i := 0; i < o.Config.Nodes; i++ {
		nodes[i] = fmt.Sprintf("swarm-node%d", i)
		resp_code, _, err := o.Validator.Get("/clients/" + nodes[i])
		if err != nil {
			return err
		}
		if resp_code == 404 {
			err = o.CreateNode(nodes[i])
			if err != nil {
				return err
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(o.Config.Nodes)
	for _, node_name := range nodes {
		go func(name string) {
			defer wg.Done()

			nd, err := node.NewNode(name, o.ServerURI, o.Config.KeyDirectory+"/"+name+".pem")
			if err != nil {
				fmt.Println(err.Error())
			}

			for i := 0; i < o.Config.Runs; i++ {
				if (i % 2) == 0 {
					err = nd.NormalRun()
				} else {
					err = nd.SearchRun()
				}
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}(node_name)
	}

	wg.Wait()

	return nil
}
