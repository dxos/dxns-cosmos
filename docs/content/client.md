---
title: Devnet Client
description: Instructions to connect to an existing `devnet` trusted node as a client.
---

Note: These are instructions to connect to an existing `devnet` trusted node as a client. To run a validator or full node, see the [setup](./validator/) doc instead.

## Endpoints

WNS

* GQL API: https://wns1.kube.moon.dxos.network/api
* GQL Console: https://wns1.kube.moon.dxos.network/console
* RPC Endpoint: tcp://wns1.kube.moon.dxos.network:26657
* Faucet GQL API: https://faucet.kube.moon.dxos.network/api

## Working with the Devnet

### Querying

To query the `devnet`, update the `wire` profile config file to the above GQL API endpoint. No other changes are required.

### Publishing

To publish records, an [account](./account/) needs to be setup.

Once the account is setup, the `wire` CLI can be used to registers records (e.g. app/bot).

Activate the `devnet` CLI profile created during account setup.

## Troubleshooting

Ensure that the CLI profile is configured correctly.

```bash
$ wire profile
$ wire profile config
```

* `services.wns.server` - must be a valid `devnet` WNS endpoint
* `services.wns.userKey` - must be the `privateKey` for the `devnet` account
* `services.wns.bondId` - must be a bond owned by the account, with sufficient funds (`wire wns bond list --owner <ACCOUNT ADDRESS>`)
