# Vault Deployment

### Purpose/Description

The vault install script will drop the configuration file for dev, test and prod environments but will be overwritten by this script.

systemd unit files that the vault install script will put in place will be overwritten with a new one from this deployment script

### Install Steps
```
unzip vault
create configuration
  HCL file
init vault
unseal vault
auth vault

consul storage backend with secret kv

write to secret (CLI/API)
create policy (CLI)
create client_token bound to the policy (CLI/API)
read secret using generated client_token (CLI/API)

```
### Provisioning
```

```
**** Coming soon ****
```

## How-Tos:

_**Bootstrapping a cluster**_

```
systemctl stop vault ;sleep 5; systemctl start vault
```

## To do

1. Sensitive info will be sourced from salt fs.  Will need to find ways to have some for of encrypted

## Bugs:
