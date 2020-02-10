package main

import (
	"bytes"
	"context"
	"errors"
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

// ContainerPID get container process id on host os
// cid: container id
func (cli *Client) ContainerPID(cid string) (string, error) {
	containerJSON, err := cli.ContainerInspect(context.Background(), cid)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(containerJSON.State.Pid), nil
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
