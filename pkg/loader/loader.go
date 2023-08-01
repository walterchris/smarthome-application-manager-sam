package loader

import "github.com/walterchris/smarthome-application-manager-sam/plugins"

var LoadFunctions []func() (plugins.Plugin, error)
