package org

import (
	"encoding/json"
	config "github.com/paulmooring/chef-swarm/lib/config"
	node "github.com/paulmooring/chef-swarm/lib/node"
	"io"
	"os"
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

	node_client, err := node.NewNode(o.Config.ValidatorName, o.ServerURI, o.Config.ValidatorKey)
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
