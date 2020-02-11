package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/skanehira/go-dockerveth"
)

func getInterfaces() (nets []net.Interface, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		if strings.Index(i.Name, "veth") != -1 {
			nets = append(nets, i)
		}
	}
	return
}

func onExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func run() error {
	cli, err := dockerveth.NewClient()
	if err != nil {
		return err
	}

	nets, err := getInterfaces()
	if err != nil {
		return err
	}

	containers, err := cli.Containers()
	if err != nil {
		return err
	}

	fmt.Println("CONTAINER\tVETH\t\tNAMES")

	for _, c := range containers {
		iflink, err := cli.ContainerIflink(c.ID)
		if err != nil {
			return err
		}

		ifindex, err := strconv.Atoi(iflink)
		if err != nil {
			return err
		}

		for _, i := range nets {
			if ifindex == i.Index {
				fmt.Printf("%s\t%s\t%s\t\n", c.ID[:12], i.Name, strings.Join(c.Names, " "))
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		onExit(err)
	}
}
