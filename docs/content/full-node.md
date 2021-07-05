---
title: Devnet Full Node Setup
description: Instructions for setting up a full node connected to the already running devnet.
---

Note: These are instructions for setting up a full node connected to the already running devnet. To run a validator from genesis, see the validator [setup](./validator/) doc.

## Requirements

* [Ubuntu server](./server/)
* [Network](./network/)

## Full Node Account Setup

Set an ENV variable with the mnemonic to be used for generating the full node account keys. Use an existing one generated earlier or create a new one using `wire keys generate`.

The mnemonic will be saved to `~/.wire/secrets` by the setup process, but also copy it to another safe location. There is no way to recover the account and associated funds if this mnemonic is lost.

```bash
$ export MNEMONIC="<MNEMONIC>"
```

Run the setup script (reset the node if prompted).

```bash
$ cd wns
$ ./scripts/setup.sh
```

## Genesis File Update

Replace `~/.wire/wnsd/config/genesis.json` file with the one in the repo.

```bash
$ cp devnet/genesis.json ~/.wire/wnsd/config/genesis.json
```

## Peer Setup

See [PEERS.md](./peers/) for the value of `persistent_peers`, and update it as described. Skip the other sections. Once peers have been setup, the node can be started.

```bash
$ ./scripts/server.sh start
```

Note: On first run, the node needs to catch up to the existing devnet block height, so the output will scroll rather quickly.

## Upgrading Full Node to Validator Node

Get funds to run a validator, then run the following command to create a validator (change the moniker below):

Note: Use the CLI passphrase from `~/.wire/secrets`.

```bash
$ wnscli tx staking create-validator \
    --moniker "<MONIKER>" \
    --chain-id "wireline" \
    --amount 10000000000000uwire \
    --pubkey $(wnsd tendermint show-validator) \
    --commission-max-change-rate "0.01" \
    --commission-max-rate "0.20" \
    --commission-rate "0.10" \
    --min-self-delegation "1" \
    --from root
```

The validator update should be visible in the wnsd logs (Note the `Updates to validators` line):

```text
I[2020-04-18|05:32:19.822] Committed state                              module=state height=227348 txs=0 appHash=9A81BD96D8098BD875FB6AA7155455A2AFAC58EBBE08C095B295916F0565438E
I[2020-04-18|05:32:25.981] Executed block                               module=state height=227349 validTxs=1 invalidTxs=0
I[2020-04-18|05:32:25.982] Updates to validators                        module=state updates=7991516DC4C16EC3E46E2CC36332F513E7B395E2:10000000
I[2020-04-18|05:32:25.987] Committed state                              module=state height=227349 txs=1 appHash=6246989C943367754D222BBBD151C0DB1745321827FE468D49B13336D9883AF4
I[2020-04-18|05:32:31.636] Executed block                               module=state height=227350 validTxs=0 invalidTxs=0
I[2020-04-18|05:32:31.641] Committed state                              module=state height=227350 txs=0 appHash=6CDE3CBBCBAC114222C54C0705FB68B54B3BBD5EF493595BA55795AD9937FFCE
```

Check that the validator is present in the validator set with non-zero voting power.

```bash
$ wnsd tendermint show-validator
cosmosvalconspub1zcjduepqe570cr3fm8scldey5d9f2paxhph5hxwr5qjzj6we6k80klmxzvgqkzj2ek

$ wnscli query tendermint-validator-set
{
  "block_height": "259388",
  "validators": [
    {
      "address": "cosmosvalcons10xg4zmwyc9hv8erw9npkxvh4z0nm890zqf9s08",
      "pub_key": "cosmosvalconspub1zcjduepqe570cr3fm8scldey5d9f2paxhph5hxwr5qjzj6we6k80klmxzvgqkzj2ek",
      "proposer_priority": "-5178125",
      "voting_power": "10000000"
    },
    {
      "address": "cosmosvalcons130c8xeen5k0k7axg787xfyw95sz4t0gjx8s4y3",
      "pub_key": "cosmosvalconspub1zcjduepqwq6xlxcxh5feuaptq3kv4wwzdln022yrw5dxfgum9zxxs0vyv9tqr475qv",
      "proposer_priority": "-10178125",
      "voting_power": "10000000"
    },
    {
      "address": "cosmosvalcons14u9ntdyydjnm7r5n9rymnhks3h0pqvcz5m008t",
      "pub_key": "cosmosvalconspub1zcjduepqt2nxz4y2zw6ffdc67wfhgrpnzqy3l64k5m40mdlsw5kw8cg7yhfqvz6a0m",
      "proposer_priority": "15534375",
      "voting_power": "9900000"
    },
    {
      "address": "cosmosvalcons1ednue6nnxyzpmj0tpa65aeedyf7gexdm00chxw",
      "pub_key": "cosmosvalconspub1zcjduepq6yqnhuytnerj7h74g87gtt3mln82h6uzv8kl4nlvlp3aps9tuthqwra85s",
      "proposer_priority": "-178125",
      "voting_power": "10000000"
    }
  ]
}
```
