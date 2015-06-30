package main

import (
	"fmt"
	config "github.com/paulmooring/chef-swarm/lib/config"
	node "github.com/paulmooring/chef-swarm/lib/node"
	"os"
)

func node_run(node_num string, chef_uri string, key_dir string) <-chan error {
	c := make(chan error)

	go func() {
		node, err := node.NewNode("test-node"+node_num, chef_uri, key_dir+"/test-node"+node_num+".pem")
		if err != nil {
			c <- err
		}

		err = node.NormalRun()
		if err != nil {
			fmt.Println(err.Error())
			c <- err
		} else {
			c <- nil
		}
	}()

	return c
}

func main() {
	cfg, err := config.LoadConfig("./config.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var org_name string
	var org_conf config.Org

	for name, conf := range cfg.Organizations {
		org_name = name
		org_conf = conf

		break
	}

	server_string := "https://" + cfg.Server + "/organizations/" + org_name

	good := 0
	bad := 0

	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 1; i <= org_conf.Runs; i++ {
		var runs []<-chan error
		if err != nil {
			fmt.Println("Error with num_nodes arg")
		}
		for i := 1; i <= org_conf.Nodes; i++ {
			c := node_run(fmt.Sprintf("%d", i), server_string, org_conf.KeyDirectory)
			runs = append(runs, c)
		}

		for _, c := range runs {
			e := <-c
			if e != nil {
				bad = bad + 1
			} else {
				good = good + 1
			}
		}
	}

	fmt.Printf("%d nodes finished\n", good)
	fmt.Printf("%d nodes failed\n", bad)
}
