package loader

import (
	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

var LoadFunctions []func(*logrus.Logger) (plugins.Plugin, error)
