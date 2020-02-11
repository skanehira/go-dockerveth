# go-dockerveth
`go-dockerveth` is CLI tool that show which docker containers are attached to which veth interfaces.  
This was inspired by [micahculpepper/dockerveth](https://github.com/micahculpepper/dockerveth).

![](https://i.imgur.com/5cMGutV.gif)

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

# If you use option `-p`, then print plane text
$ dockerveth -p
ad448862880e    veth393c35f     vimgolf         ruby            bash
eadea1368853    veth43f9468     golang1.13      golang:latest   bash

# help
$ dockerveth -h
Usage of dockerveth:
  -p    make plane text(default is make table)
```

## Author
skanehira
