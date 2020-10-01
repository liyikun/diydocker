package main

import (
	"github.com/liyikun/diydocker/container"
	"github.com/sirupsen/logrus"
	"os"
)

func Run(tty bool, command string)  {
	parent := container.NewParentProcess(tty, command)

	if err := parent.Start() ; err != nil {
		logrus.Error(err)
	}

	parent.Wait()

	os.Exit(-1)
}