// +build linux

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)


const usage = "simple docker"


func main()  {
	app := cli.NewApp()
	app.Name = "lykdocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})

		logrus.SetOutput(os.Stdout)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}