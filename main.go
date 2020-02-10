package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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

func main() {
	nets, err := getInterfaces()
	if err != nil {
		onExit(err)
	}

	containers()

	fmt.Println(nets)
}
