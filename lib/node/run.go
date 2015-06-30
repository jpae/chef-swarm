package node

import (
	"github.com/nu7hatch/gouuid"
)

func (nd Node) dataSearch() error {
	_, _, err := nd.Get("/search/test?q=id:*")
	if err != nil {
		return err
	}

	return nil
}

func (nd Node) nodeSearch() error {
	_, _, err := nd.Get("/search/node?q=id:*")
	if err != nil {
		return err
	}

	return nil
}

func (nd Node) getDataBag(bag_name string, item_name string) error {
	_, _, err := nd.Get("/data/" + bag_name + "/" + item_name)
	if err != nil {
		return err
	}

	return nil
}

func (nd Node) runBody(run_type string) error {
	switch run_type {
	case "normal":
		nd.dataSearch()
		sleep_splay(1200)
		nd.nodeSearch()
		sleep_splay(1000)
		nd.getDataBag("test", "item1")
		sleep_splay(750)
		nd.getDataBag("test", "item15")
		sleep_splay(5000)
	default:
		nd.Get("/nodes/" + nd.Name)
	}

	return nil
}

func (nd Node) runWrapper(run_type string) error {
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	run_uuid := u.String()
	start_time := timestamp()

	_, _, err = nd.Get("/nodes/" + nd.Name)
	if err != nil {
		return err
	}
	sleep_splay(100)

	_, _, err = nd.Get("/environments/_default")
	if err != nil {
		return err
	}
	sleep_splay(100)

	body := start_run_body(run_uuid, start_time)
	_, _, err = nd.Post("/reports/nodes/"+nd.Name+"/runs", body)
	if err != nil {
		return err
	}
	sleep_splay(50)

	body = cookbook_versions_body()
	_, _, err = nd.Post("/environments/_default/cookbook_versions", body)
	if err != nil {
		return err
	}
	sleep_splay(1000)

	err = nd.runBody(run_type)
	if err != nil {
		return err
	}
	sleep_splay(1000)

	body = node_body(nd.Name)
	_, _, err = nd.Put("/nodes/"+nd.Name, body)
	if err != nil {
		return err
	}
	sleep_splay(200)

	body = end_run_body(run_uuid, start_time)
	_, _, err = nd.Post("/reports/nodes/"+nd.Name+"/runs/"+run_uuid, body)
	if err != nil {
		return err
	}

	return nil
}

func (nd Node) EmptyRun() error {
	err := nd.runWrapper("empty")
	if err != nil {
		return err
	}

	return nil
}

func (nd Node) NormalRun() error {
	err := nd.runWrapper("normal")
	if err != nil {
		return err
	}

	return nil
}
