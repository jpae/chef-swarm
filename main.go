package main

import (
	"fmt"
	node "github.com/paulmooring/chef-swarm/lib/node"
	"os"
	"strconv"
)

func node_run(node_num string) <-chan error {
	c := make(chan error)

	go func() {
		node, err := node.NewNode("test-node"+node_num, os.Args[1], "examples/nodes/test-node"+node_num+".pem")
		if err != nil {
			c <- err
		}

		err = node.NormalRun()
		if err != nil {
			c <- err
		} else {
			c <- nil
		}
	}()

	return c
}

func main() {
	good := 0
	bad := 0

	num_ccrs, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 1; i <= num_ccrs; i++ {
		var runs []<-chan error
		num_nodes, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error with num_nodes arg")
		}
		for i := 1; i <= num_nodes; i++ {
			c := node_run(fmt.Sprintf("%d", i))
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
