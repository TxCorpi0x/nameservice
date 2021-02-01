# nameservice

**nameservice** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Prepare development environment on Docker

This repo is ready to use by Visual Studio Code and the Docker. for extreme simplicity issue these commands:

```
git clone github.com/vjdmhd/nameservice 
cd nameservice
code .
```
VSCode can create your container without getting your hands dirty.
after that if permission related problem raised issue this in docker terminal:
```
chown -R vscode /go/src/github.com/vjdmhd/nameservice
```
Cheers! you are ready to go.

## Prepare development environment on Manually

Please follow steps mentioned in this tutorial:
[Cosmos Network Tutorials](https://tutorials.cosmos.network/nameservice/tutorial/00-intro.html)

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Testing
To use testing functions issue this command at the root of repo:
```
go test -v
```

## Version compatibility
This repo is on top of Cosmos-SDK@0.39.1, if you want to modify code base, keep in mind that importing new module versions of Cosmos-SDK might destroy your environment.



## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
