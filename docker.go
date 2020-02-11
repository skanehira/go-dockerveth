package dockerveth

import (
	"bytes"
	"context"
	"errors"
	"net"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

var (
	ErrEmptyExecID = errors.New("empty exec id")
)

// Client wrap docker client
type Client struct {
	*client.Client
}

// NewClient create new docker client
func NewClient() (*Client, error) {
	c, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

// Containers get contianers
func (cli *Client) Containers() ([]types.Container, error) {
	return cli.ContainerList(context.Background(), types.ContainerListOptions{})
}

// ContainerIflink get container iflink
func (cli *Client) ContainerIflink(cid string) (string, error) {
	res, err := cli.ContainerExecCreate(context.Background(), cid, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"cat", "/sys/class/net/eth0/iflink"},
	})

	if err != nil {
		return "", err
	}
	if res.ID == "" {
		return "", ErrEmptyExecID
	}

	hijackResponse, err := cli.ContainerExecAttach(context.Background(), res.ID, types.ExecConfig{})
	if err != nil {
		return "", err
	}
	defer hijackResponse.Close()

	var dstout, dsterr = &bytes.Buffer{}, &bytes.Buffer{}
	_, err = stdcopy.StdCopy(dstout, dsterr, hijackResponse.Reader)
	if err != nil {
		return "", err
	}

	// if exec command return error
	if dsterr.Len() != 0 {
		return "", errors.New(dsterr.String())
	}

	return strings.Trim(dstout.String(), "\n"), nil
}

// GetContainerInfo get container info used by output
func (cli *Client) GetContainerInfo() ([][]string, error) {
	containers, err := cli.Containers()
	if err != nil {
		return nil, err
	}

	nets, err := getIPInterfaces()
	if err != nil {
		return nil, err
	}

	var rows [][]string
	for _, c := range containers {
		iflink, err := cli.ContainerIflink(c.ID)
		if err != nil {
			return nil, err
		}

		ifindex, err := strconv.Atoi(iflink)
		if err != nil {
			return nil, err
		}

		for _, i := range nets {
			if ifindex == i.Index {
				row := []string{
					c.ID[:12],
					i.Name,
					parseContainerNames(c.Names),
					c.Image,
					c.Command,
				}
				rows = append(rows, row)
			}
		}
	}
	return rows, nil
}

// getIPInterfaces get ip info
func getIPInterfaces() (nets []net.Interface, err error) {
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

func parseContainerNames(names []string) string {
	var s []string
	for _, n := range names {
		s = append(s, n[1:])
	}
	return strings.Join(s, " ")
}
