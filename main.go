package main

import (
	"fmt"
	config "github.com/paulmooring/chef-swarm/lib/config"
	org "github.com/paulmooring/chef-swarm/lib/org"
	"os"
)

func main() {
	cfg, err := config.LoadConfig("./config.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var org_name string
	var org_conf config.OrgConfig

	for name, conf := range cfg.Organizations {
		org_name = name
		org_conf = conf

		break
	}

	server_string := "https://" + cfg.Server + "/organizations/" + org_name

	o, err := org.NewOrg(server_string, org_conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = o.ChefRuns()
	if err != nil {
		fmt.Println(err.Error())
	}
}
