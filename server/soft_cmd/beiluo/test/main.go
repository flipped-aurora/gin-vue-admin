package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logger.Setup()
	command := os.Args[1]
	args := os.Args[2]
	logrus.Infof("command name: %s", command)
	logrus.Infof("args name: %s", args)
}
