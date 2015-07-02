# A Load test for chef-server

This is very much a work in progress, everything below will more than likely change at some point in the near future.

## Set up

`chef-swarm` will create clients/nodes as needed, but it expects a validator key to be available.  Additionally cookbooks should be uploaded and a `test` data bag should be created.  Runs will still work without cookbooks or data bags, but they would be significantly less meaningful and return a lot of HTTP 412/404 errors.

## Config

Config is read from a toml file in the directory the program is run from.  Currently `./config.toml` is hardcoded, this will/should be fixed.  There's an example config included in this repo.

