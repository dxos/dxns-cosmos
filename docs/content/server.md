---
title: Ubuntu Setup
description: Instructions to setup a Ubuntu node.
---

## Pre-Setup

Install Ubuntu Server 18.04 LTS on the xbox by booting via a USB stick.

https://ubuntu.com/download/server

Instructions below assume:

* Terminal based remote login (SSH) to the xbox with a user that has `sudo` access, setup during OS installation.
* GitHub [ssh agent forwarding](https://developer.github.com/v3/guides/using-ssh-agent-forwarding/) setup to clone repositories.

## Setup

Update packages.

```bash
$ sudo apt-get update && sudo apt-get upgrade -y
```

Install essential development packages and utils.

```bash
$ sudo apt-get install build-essential autossh jq -y
```

Download and install golang binaries.

```bash
$ wget https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.13.8.linux-amd64.tar.gz
```

Add `/usr/local/go/bin` to the `PATH` environment variable by editing `~/.profile`.

```
export PATH=$PATH:/usr/local/go/bin
```

Apply the changes to `~/.profile`.

```bash
$ source ~/.profile
```

Test that golang has been successfully installed on the machine.

```bash
$ go version
go version go1.13.8 linux/amd64
```

Setup `gomod` ENV vars.

```bash
$ mkdir -p $HOME/go/bin
$ echo "export GOPATH=$HOME/go" >> ~/.profile
$ echo "export GOBIN=\$GOPATH/bin" >> ~/.profile
$ echo "export PATH=\$PATH:\$GOBIN" >> ~/.profile
$ echo "export GO111MODULE=on" >> ~/.profile
$ source ~/.profile
```

Clone the repo then build and install the binaries.

```bash
$ git clone git@github.com:wirelineio/wns.git
$ cd wns
$ make install
```

Test that the following commands work.

```bash
$ wnsd help
$ wnscli help
```

Setup `nvm` to install Node.js (https://github.com/nvm-sh/nvm).

```bash
$ cd ~
$ curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.2/install.sh | bash
```

Logout and re-login to the xbox.

Test `nvm` is installed.

```bash
$ nvm help
```

Install Node.js.

```bash
$ nvm install v12.16.1
```

Test that Node.js is installed.

```bash
$ node -v
v12.16.1
```

Install `yarn`.

```bash
$ npm install -g yarn
```

Test that `yarn` is installed.

```bash
$ yarn --version
1.22.0
```

Edit `~/.profile` and add `yarn` bin folder to the `PATH`.

```
export PATH=$PATH:$HOME/.yarn/bin
```

Apply the changes to `~/.profile`.

```bash
$ source ~/.profile
```

Install the `wire` CLI.

```bash
$ yarn global add @dxos/cli@beta
```

Check version of `wire` CLI.

```bash
$ wire version
v1.0.0-beta.13
```
