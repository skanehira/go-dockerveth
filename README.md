# go-dockerveth
`go-dockerveth` is CLI tool that show which docker containers are attached to which veth interfaces.  
This is inspired by [micahculpepper/dockerveth](https://github.com/micahculpepper/dockerveth).

![](https://i.imgur.com/5cMGutV.gif)

## Features
- Doesn't need sudo
- Doesn't need docker cli
- Can be used in combination with other commands in plain text output

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
+--------------+-------------+------------+---------------+------+
|  CONTAINER   |    VETH     |   NAMES    |     IMAGE     | CMD  |
+--------------+-------------+------------+---------------+------+
| ad448862880e | veth393c35f | vimgolf    | ruby          | bash |
| eadea1368853 | veth43f9468 | golang1.13 | golang:latest | bash |
+--------------+-------------+------------+---------------+------+

# If you use option `-p`, then print plain text
$ dockerveth -p
ad448862880e    veth393c35f     vimgolf         ruby            bash
eadea1368853    veth43f9468     golang1.13      golang:latest   bash

# help
$ dockerveth -h
Usage of dockerveth:
  -p    make plain text(default is make table)
```

If you want to capture packet from container, you cat do it following.
```sh
# fzf is fuzzy finder, you can install it from https://github.com/junegunn/fzf
# termshark is tui for wireshark, you can install if from https://github.com/gcla/termshark
sudo termshark -i $(dockerveth -p | fzf | awk "{print \$2}")
```

## Author
skanehira
