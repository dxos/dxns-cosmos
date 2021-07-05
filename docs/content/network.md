---
title: Network
description: Intructions for setting up connectivity to devnet nodes.
---

# Network

Nodes need to be able to connect to each other for the devnet to work.

Ideally, each node should be provisioned with a static IPv4 address.

If a node doesn't have a static IP, use remote port forwarding or a VPN to setup network connectivity.

## Remote Port Forwarding

Setup a SSH tunnel (remote port forwarding) between your devnet node and a remote host.

Note: See https://www.digitalocean.com/community/tutorials/how-to-configure-custom-connection-options-for-your-ssh-client for configuring SSH client.

In the example below, `user@host` is the remote user and hostname/IP configured for public key based SSH login and which is reachable from other nodes. Prior to starting the tunnel, check that `ssh <user@host>` works, to verify `~/.ssh/config` settings.

To start the tunnel.

```bash
$ cd wns
$ ./scripts/tunnel.sh start <user@host>
```

To view the tunnel logs.

```bash
$ ./scripts/tunnel.sh log
```

To stop the tunnel.

```bash
$ ./scripts/tunnel.sh stop
```

## References

* https://www.ssh.com/ssh/tunneling/example#remote-forwarding
* https://www.everythingcli.org/ssh-tunnelling-for-fun-and-profit-autossh
