package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/skanehira/go-dockerveth"
)

var (
	isPlain = flag.Bool("p", false, "make plain text(default is make table)")
)

func onExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func makeTable(rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{"CONTAINER", "VETH", "NAMES", "IMAGE", "CMD"}
	table.SetHeader(header)
	headerColor := tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor}
	var headerColors []tablewriter.Colors
	for i := 0; i < len(header); i++ {
		headerColors = append(headerColors, headerColor)
	}
	table.SetHeaderColor(headerColors...)
	table.AppendBulk(rows)
	table.Render()
}

func makePlainText(rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(rows)
	table.Render()
}

func run() error {
	flag.Parse()
	cli, err := dockerveth.NewClient()
	if err != nil {
		return err
	}

	rows, err := cli.GetContainerInfo()
	if err != nil {
		return err
	}

	if *isPlain {
		makePlainText(rows)
	} else {
		makeTable(rows)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		onExit(err)
	}
}
