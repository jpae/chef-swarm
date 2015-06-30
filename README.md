# A Load test for chef-server

This is very much a work in progress, everything below will more than likely change at some point in the near future.

## Set up

Currently a client/node name format of `test-nodeX` is expected:

```
# Create the clients
for i in  {1..100} ; do
    knife client create -d -u example-org-validator -k .chef/example-org-validator.pem test-node${i} > nodes/test-node${i}.pem
done

# Create the nodes
for i in  {1..100} ; do
    knife node create -d -u test-node${i} -k nodes/test-node${i}.pem test-node${i}
done
```

## Config

Config is read from a toml file in the directory the program is run from.  Currently `./config.toml` is hardcoded, this will/should be fixed.  There's an example config included in this repo.

## Example invocations

chef-swarm
