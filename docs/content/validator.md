---
title: Devnet Validator Node Setup
description: Instructions for setting up a validator node from genesis.
---

Note: These are instructions for setting up a validator node from genesis. To run a full node connected to the already running devnet, see the full node [setup](./full-node/) doc.

Note: See the [client doc](./client/) for instructions to connect to an existing trusted node.

## Requirements

* [Ubuntu server](./server/)
* [Network](./network/)

## Validator Account Setup

Note: Run this step on every validator node.

Set an ENV variable with the mnemonic to be used for generating the validator account keys. Use an existing one generated earlier or create a new one using `wire keys generate`.

The mnemonic will be saved to `~/.wire/secrets` by the setup process, but also copy it to another safe location. There is no way to recover the account and associated funds if this mnemonic is lost.

```bash
$ export MNEMONIC="<MNEMONIC>"
```

Run the setup script (reset the node if prompted).

```bash
$ cd wns
$ ./scripts/setup.sh
```

Check-in the genesis transaction file created in `~/.wire/wnsd/config/gentx` to the `wns/devnet/gentx` folder.

Get the validator account address.

```bash
$ wnscli keys show root -a
cosmos174hrcf4x9nhwzt82qwns65esa0a7u05425jftp
```

Update SEED_ACCOUNTS.md with a new entry (validator address as above):

Note: Do NOT run this command, only copy it to the above file.

```text
wnsd add-genesis-account cosmos174hrcf4x9nhwzt82qwns65esa0a7u05425jftp 100000000000000uwire
```

## Genesis File Generation

Note: Run this step only on the initial validator node, to generate the consolidated `genesis.json` file.

Run the above setup. Delete existing contents in `~/.wire/wnsd/config/gentx` folder and copy all the gentx files from the repo to `~/.wire/wnsd/config/gentx`.

```bash
$ rm ~/.wire/wnsd/config/gentx/*
$ cp devnet/gentx/* ~/.wire/wnsd/config/gentx
```

Add the genesis accounts from [SEED_ACCOUNTS.md](./seed-accounts/).

Re-generate the genesis.json file.

```bash
$ wnsd collect-gentxs
$ wnsd validate-genesis
```

Check-in the updated `~/.wire/wnsd/config/genesis.json` file to `wns/devnet/genesis.json`.

```bash
$ cp ~/.wire/wnsd/config/genesis.json devnet/genesis.json
```

## Genesis File Update

Note: Run this step on every validator node.

All validators should replace their `~/.wire/wnsd/config/genesis.json` file with the one in the repo, after the consolidated genesis.json has been generated.

```bash
$ cp devnet/genesis.json ~/.wire/wnsd/config/genesis.json
```

## Peer Setup

Note: Run this step on every validator node.

See [PEERS.md](./peers/) to configure your node with peers. Once peers have been setup, the node can be started.

```bash
$ ./scripts/server.sh start
```

The devnet will generate blocks once 2/3 of voting power is online.

Check WNS logs to verify blocks are being generated.

```bash
$ ./scripts/server.sh log
```

## Post Setup

Note: To be run on each node once the devnet is operational and blocks are being generated.

Update `~/.profile` with account private key and node address. The private key can be looked up in `~/.wire/secrets`.

```
export WIRE_WNS_ENDPOINT="http://localhost:9473/graphql"
export WIRE_WNS_USER_KEY="<PRIVATE KEY>"
```

Apply the changes to `~/.profile`.

```bash
$ source ~/.profile
```

Generate a bond ID, which is required to pay for records.

```bash
$ wire wns create-bond --type uwire --quantity 10000000000
{
    "submit": "9F3E05DECE29D1B20F8148B8AEDA31058094036C2971ACA963A6ABE83A59587E"
}
```

List the bond IDs by owner address (`~/.wire/secrets` has the address to use).

```bash
$ wire wns list-bonds --owner cosmos1np8f3zzu6xss0m2rh2k7ugawegw0x29gh9n2lq
[
    {
        "id": "36550248ce6bd1d391825bc9111956dd899ac3ca03c238a20f79be49c8a9f806",
        "owner": "cosmos1np8f3zzu6xss0m2rh2k7ugawegw0x29gh9n2lq",
        "balance": [
            {
                "type": "uwire",
                "quantity": "10000000000"
            }
        ]
    }
]
```

Update `~/.profile` with the bond ID.

```
export WIRE_WNS_BOND_ID="<BOND ID>"
```

Apply the changes to `~/.profile`.

```bash
$ source ~/.profile
```

## Unjail Validator

If a validator node is offline for too long, it will get jailed.

Check the status of a validator.

```bash
$ wnsd tendermint show-validator
cosmosvalconspub1zcjduepqt2nxz4y2zw6ffdc67wfhgrpnzqy3l64k5m40mdlsw5kw8cg7yhfqvz6a0m

$ wnscli query staking validators
```

Note the `"jailed": true` in the output for the corresponding validator (`consensus_pubkey`).

```text
  {
    "operator_address": "cosmosvaloper174hrcf4x9nhwzt82qwns65esa0a7u0540qxu8j",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqt2nxz4y2zw6ffdc67wfhgrpnzqy3l64k5m40mdlsw5kw8cg7yhfqvz6a0m",
    "jailed": true,
    "status": 1,
    "tokens": "9900000000000",
    "delegator_shares": "10000000000000.000000000000000000",
    "description": {
      "moniker": "xbox-nuc",
      "identity": "",
      "website": "",
      "details": ""
    },
    "unbonding_height": "227413",
    "unbonding_time": "2020-05-09T05:39:39.645217342Z",
    "commission": {
      "commission_rates": {
        "rate": "0.100000000000000000",
        "max_rate": "0.200000000000000000",
        "max_change_rate": "0.010000000000000000"
      },
      "update_time": "2020-04-02T13:13:39.434616781Z"
    },
    "min_self_delegation": "1"
  }
```

Create a transaction to unjail the validator.

Note: Use the CLI passphrase from `~/.wire/secrets`.

```bash
$ wnscli tx slashing unjail --from root
```

The validator set change should be visible in the logs:

```text
I[2020-04-20|05:12:25.298] Committed state                              module=state height=256451 txs=0 appHash=063AF2F36CE278B7C7D54B940DA8B0E2704906D04E3C6CBC818B1C45A74CDF1B
I[2020-04-20|05:12:31.099] Executed block                               module=state height=256452 validTxs=1 invalidTxs=0
I[2020-04-20|05:12:31.100] Updates to validators                        module=state updates=AF0B35B4846CA7BF0E9328C9B9DED08DDE103302:9900000
I[2020-04-20|05:12:31.101] Committed state                              module=state height=256452 txs=1 appHash=F827190E5136B5D79898A7CBAD89BB12EC963DF3E8E5A2BF9C298B34C8A9B2EE
I[2020-04-20|05:12:37.083] Executed block                               module=state height=256453 validTxs=0 invalidTxs=0
```

Check that the validator is unjailed (grep should return a matching line).

```bash
$ wnscli query tendermint-validator-set | grep "$(wnsd tendermint show-validator)"
```

Check voting power (should be non-zero):

```
$ wnscli status
{
  "node_info": {
    "protocol_version": {
      "p2p": "7",
      "block": "10",
      "app": "0"
    },
    "id": "02a9e173f4bc1c3b9969d87aa2c1899ffab60901",
    "listen_addr": "tcp://0.0.0.0:26656",
    "network": "wireline",
    "version": "0.32.2",
    "channels": "4020212223303800",
    "moniker": "xbox-nuc",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://127.0.0.1:26657"
    }
  },
  "sync_info": {
    "latest_block_hash": "8D0FF6DE0E06E6E0B4C8AEB203002CA29D0F00F1055E5DE365BC8523041EF233",
    "latest_app_hash": "503B67B7BA15F75F1FF2E1F97501EF7EED9E60BD072970F4CBFA48D245398B31",
    "latest_block_height": "256487",
    "latest_block_time": "2020-04-20T05:15:49.667060991Z",
    "catching_up": false
  },
  "validator_info": {
    "address": "AF0B35B4846CA7BF0E9328C9B9DED08DDE103302",
    "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "WqZhVIoTtJS3GvOTdAwzEAkf6ram6v238HUs4+EeJdI="
    },
    "voting_power": "9900000"
  }
}
```
