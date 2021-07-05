---
title: Peers
description: Devnet peers.
---

Update `~/.wire/wnsd/config/config.toml` with:

```text
persistent_peers = "13578da4407a6884c765367769d6d7dfe9a20923@139.178.68.130:26656,c48b4f6f00538dee0ac268c2ae20ec67ec57b771@139.178.68.131:26656,02a9e173f4bc1c3b9969d87aa2c1899ffab60901@wns1.deepstacksoft.com:26656"
```

## Adding Peers

Get the Tendermint Node ID of the new peer:

```bash
$ wnsd tendermint show-node-id
```

Get the public hostname/IP for the machine to add as a new peer and update (`<node-id>@<host/IP>:26656`) the above list. Peers are separated by commas.


## Troubleshooting

* If the node hostname/IP is not routable, nodes might have trouble connecting to each other. Try setting `addr_book_strict = false` in `~/.wire/wnsd/config/config.toml`.
* If the node does not have a static IP, [reverse port forwarding](./network/) can be used to tunnel through a remote machine that has a public IP/hostname.
