package main

import (
	"github.com/Sirupsen/logrus"
)

func main() {
	var f logrus.JSONFormatter
	logrus.WithFields(logrus.Fields{})
	_ = f
}
