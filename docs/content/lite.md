---
title: Devnet Lite Node Setup
description: Instructions for setting up a lite node connected to the already running devnet.
---

# WNS Lite

WNS Lite is a light weight read-only cache of the WNS graph database.

* Choose a trusted full-node for initialization of the lite node
* Follow the lite node setup instructions below

## Getting Started

### Installation

Install WNS, then test that the following command works:

```bash
$ wnsd-lite help
```

### Initializing the Lite Node

The lite node needs an upstream full-node to function. It connects to the `cosmos-sdk`/`tendermint` RPC endpoint on the full-node to pull state, with merkle proofs.

To expose the RPC endpoint on a full-node outside the local host, edit `~/.wire/wnsd/config/config.toml` and set the `laddr` in the `[rpc]` section to a static IP or `0.0.0.0`. Ensure the RPC port (`26657` by default) isn't firewalled.

For example:

```text
##### rpc server configuration options #####
[rpc]

# TCP or UNIX socket address for the RPC server to listen on
laddr = "tcp://0.0.0.0:26657"
```

Initialize the lite node:

```bash
$ ./scripts/lite/setup.sh --node "<WNS RPC ENDPOINT>"
```

Example:

```bash
$ ./scripts/lite/setup.sh --node "tcp://wns1.kube.moon.dxos.network:26657"
```

## Working with the Lite Node

Start the node:

```bash
$ ./scripts/lite/server.sh start --node "<WNS RPC ENDPOINT>"
```

Example:

```bash
$ ./scripts/lite/server.sh start --node "tcp://wns1.kube.moon.dxos.network:26657"
```

To enable the lite node to periodically discover additional RPC endpoints from WNS, pass it a GQL API endpoint (in the example below, it's the lite node endpoint itself).

```bash
$ ./scripts/lite/server.sh start --node "tcp://wns1.kube.moon.dxos.network:26657" --endpoint "http://127.0.0.1:9475/api"
```

Optionally, pass `--reset` to also do a full reset before starting the node.

Example:

```bash
$ ./scripts/lite/server.sh start --node "tcp://wns1.kube.moon.dxos.network:26657" --endpoint "http://127.0.0.1:9475/api" --reset
```

Test if the node is up:

```bash
$ ./scripts/lite/server.sh test
```

View the logs:

```bash
$ ./scripts/lite/server.sh log
```

Stop the node:

```bash
$ ./scripts/lite/server.sh stop
```

### RPC Endpoint Discovery

Currently, RPC endpoints are discovered by querying for `kube` type records with a `wns.rpc` field.

To register a `kube` with a WNS RPC endpoint:

```bash
$ wire kube register --id 'wrn:kube:ashwinp/wns1' --version 0.0.1 --data.wns.rpc='tcp://45.79.120.249:26657'
```

```
$ wire wns record get --id Qmae4rq7QzLwz4qrqoDHa29w3CXQGSA8G766zLmM5yWVrU
[
  {
    "id": "Qmae4rq7QzLwz4qrqoDHa29w3CXQGSA8G766zLmM5yWVrU",
    "type": "wrn:kube",
    "name": "ashwinp/wns1",
    "version": "0.0.1",
    "owners": [
      "233b436a205539f0f8082507e300fc5f3ca9eb0a"
    ],
    "bondId": "8a359128068c85f9982a36308772057d098f16dc21288e312205bdf60a6961e9",
    "createTime": "2020-04-22T07:58:31.839889941",
    "expiryTime": "2021-04-22T07:58:31.839889941",
    "attributes": {
      "type": "wrn:kube",
      "version": "0.0.1",
      "wns": "{\"rpc\":\"tcp://45.79.120.249:26657\"}",
      "name": "ashwinp/wns1"
    }
  }
]
```
