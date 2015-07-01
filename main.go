package main

import (
	"fmt"
	config "github.com/paulmooring/chef-swarm/lib/config"
	org "github.com/paulmooring/chef-swarm/lib/org"
	"os"
	"sync"
)

func main() {
	cfg, err := config.LoadConfig("./config.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var wait_group sync.WaitGroup
	wait_group.Add(len(cfg.Organizations))

	for name, conf := range cfg.Organizations {
		go func(org_name string, org_conf config.OrgConfig) {
			defer wait_group.Done()

			server_string := "https://" + cfg.Server + "/organizations/" + org_name

			o, err := org.NewOrg(server_string, org_conf)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = o.ChefRuns()
			if err != nil {
				fmt.Println(err.Error())
			}
		}(name, conf)
	}

	wait_group.Wait()
}
