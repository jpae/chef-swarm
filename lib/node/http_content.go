package node

import (
	"bytes"
)

func node_body(name string) *bytes.Buffer {
	body := `{
  "name":"` + name + `",
  "chef_environment": "_default",
  "chef_type": "node",
  "automatic": {},
  "normal":{},
  "default":{},
  "override":{},
  "json_class": "Chef::Node",
  "run_list": []
}`

	return string_to_reader(body)
}

func end_run_body(run_uuid string, start_time string) *bytes.Buffer {
	body := `{
  "action": "end",
  "resources": [
    {
      "type": "execute",
      "name": "an execute resource",
      "id": "echo 'hello'",
      "after": {},
      "before": {},
      "duration": "15340",
      "delta": "",
      "result": "run",
      "cookbook_name": "dotfiles",
      "cookbook_version": "1.1.1"
    }
  ],
  "status": "success",
  "run_list": "[\"role[base]\"]",
  "start_time": "` + start_time + `",
  "end_time": "` + timestamp() + `",
  "total_res_count": "46",
  "data": {}
}`

	return string_to_reader(body)
}

func start_run_body(run_uuid string, start_time string) *bytes.Buffer {
	body := `{
  "action": "start",
  "run_id": "` + run_uuid + `",
  "start_time": "` + start_time + `"
}`

	return string_to_reader(body)
}

func cookbook_versions_body() *bytes.Buffer {
	body := `{
  "run_list": [
    "apt",
    "build-essential",
    "chef-client@4.3.0",
    "cron",
    "dotfiles",
    "ntp@1.8.6",
    "users@1.8.2",
    "windows",
    "yum",
    "yum-epel"
  ]
}`

	return string_to_reader(body)
}
