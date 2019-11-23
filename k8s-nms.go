package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/smarkm/k8s-nms/cmd"
)

//Common var
var (
	Version string = "0.0.1"
)

func main() {

	ui := &cli.BasicUi{Reader: os.Stdin, Writer: os.Stdout, ErrorWriter: os.Stderr}
	c := cli.NewCLI("k8s-nms", Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"web": func() (cli.Command, error) {
			return &cmd.StartApp{UI: ui}, nil
		},
	}
	_, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(0)
}
