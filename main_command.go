package main

import (
	"github.com/liyikun/diydocker/container"
	"github.com/urfave/cli"
	"fmt"
	"github.com/sirupsen/logrus"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: "Create a container with namespace and cgroups limit simpledocker run -ti [command]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")

		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name: "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		cmd := context.Args().Get(0)
		logrus.Infof("command %s", cmd)

		err := container.RunContainerInitProcess(cmd, nil)

		return err
	},
}
