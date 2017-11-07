# Vault Deployment

## Purpose/Description

The vault install walks through the steps require to install, configure, initialize a vault server. 
After the initial steps, it covers how to write, create policy and read from secret backend.

Consul is the configured storage backend. It proviced auto health check.

```
## Install Steps
```

unzip vault

create configuration
  HCL file - vault-server.hcl

init vault
  Store the generated keys and master vault token

vault unseal
  Use the generated keys

write to secret (CLI/API)

create policy (CLI)

create client_token bound to the policy (CLI/API)

read secret using generated client_token (CLI/API)

```
## Provisioning
```

```
**** Coming soon ****
```

## How-Tos:
vault unseal

vault policy-write <name_of_policy> policy.hcl

vault token-create -policy=<name_of_policy> -format=json |jq -r .auth.client_token

vault write secret/path-to-key value=password

vault read secret/path-to-key

vault token-revoke <client_token>


_**Bootstrapping a cluster**_

```
systemctl stop vault ;sleep 5; systemctl start vault
```

## To do

1. Sensitive info will be sourced from salt fs.  Will need to find ways to have some for of encrypted

## Bugs:
