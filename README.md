# BDEVault

BDEVault is a tool intended to make your life easier, in regards to BitLocker.
It submits the local computer's BitLocker recovery keys into a [Hashicorp Vault](https://vaultproject.io/) of your choosing.

## Build

```
go build ./cmd/bdevault
```

## Usage

```
pandentia@Tali ~/d/bdevault (main)> ./bdevault.exe -h
Usage of bdevault.exe:
  -address string
        Vault address
  -key string
        Vault key to write into (default "tali")
  -mount string
        Vault mount to write to (default "secrets")
  -prefix string
        Vault mount keyspace prefix (default "bitlocker/")
  -token string
        Vault token
```

In effect, give it a Vault address, a token, and a mount path. The former two will also be pulled from `VAULT_ADDR` and `VAULT_TOKEN`, respectively.

The Vault key default is derived from system hostname, converted to lowercase.

## Caveats

- Must be run as admin for `manage-bde` to work
- Incompatible with Vault KV1
- Incompatible with KV2's check-and-set parameter
- Always overwrites the key, regardless of content. It is therefore inadvisable to run this as a scheduled task.
