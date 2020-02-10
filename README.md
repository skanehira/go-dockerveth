# go-dockerveth
`go-dockerveth` is CLI tool that show which docker containers are attached to which veth interfaces.
This was inspired by [micahculpepper/dockerveth](https://github.com/micahculpepper/dockerveth).

## Requirements
- Go 1.13.4 or higher

## Support OS
- Linux

## Installation
```sh
$ git clone https://github.com/skanehira/go-dockerveth
$ cmd go-dockerveth/cmd/dockerveth
$ go install
```

## Usage
```sh
$ dockerveth
CONTAINER       VETH            NAMES
2c5c7a5c1804    veth1ce36c6     /php
c501ed5d2dee    vethfafa2ae     /golang
```

## Author
skanehira
