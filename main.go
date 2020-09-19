// +build linux

package main

import "github.com/urfave/cli"


const usage = "simple docker"

func main()  {
	app := cli.NewApp()
	app.Name = "lykdocker"
	app.Usage = usage

	app.Commands = []cli.Command()
}