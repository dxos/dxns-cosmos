---
title: Devnet 2 (Moon Release)
description: Instructions for setting up a validator at genesis.
---

## Setup

Build and install the binaries:

```bash
$ cd wns
$ git checkout release-moon
$ make install
```

## Validator at Genesis

Choose a name for your validator (short alphanumeric string, without spaces):

```bash
$ export VALIDATOR_NAME='<NAME>'
```

Generate the node and validator keys and save the generated mnemonic to a safe place:

```bash
$ ./scripts/setup.sh --reset --chain-id devnet-2 --node-name $VALIDATOR_NAME
$ cp networks/devnet-2/genesis.json ~/.wire/wnsd/config/genesis.json
$ wnscli keys add $VALIDATOR_NAME
```

Generate a staking transaction (sign with your validator key):

```bash
$ wnsd add-genesis-account $(wnscli keys show $VALIDATOR_NAME -a) 100000000000000uwire
$ wnsd gentx --name $VALIDATOR_NAME --amount 10000000000000uwire --output-document $VALIDATOR_NAME.json
```

Commit the generated file to `wns/networks/devnet-2/gentx` folder and push the changes to `release-moon` branch.

## Starting the Chain

### Genesis

Copy over the final genesis.json:

```bash
$ cp networks/devnet-2/genesis.json ~/.wire/wnsd/config/genesis.json
```

### Peers

Update `~/.wire/wnsd/config/config.toml` with:

```text
persistent_peers = "b0e7891bd855a2e4c52104cb194456bbc5bd2922@dxns1.devnet-2.deepstacksoft.com:26656,4661cd608d372ab887b69960d2c850f63d4136ac@wire0.vaasl.io:26656,0d103010b049ca9fd675878390289dffc068f4d1@134.122.84.173:26656"
```

To list your validator node as a peer for others, first get the node ID:

```bash
$ wnsd tendermint show-node-id
```

Then, get the public hostname/IP for the machine and update (`<node-id>@<host/IP>:26656`) the above list. Peers are separated by commas.

### Run

```bash
$ ./scripts/server start --tail
```

## Synchronized Chain Reset

For Moon milestone:

```bash
$ cd wns
$ git pull
$ git checkout release-moon
$ make install
$ cp networks/devnet-2/genesis.json ~/.wire/wnsd/config/genesis.json
$ wnsd unsafe-reset-all
$ ./scripts/server.sh start --tail
```
