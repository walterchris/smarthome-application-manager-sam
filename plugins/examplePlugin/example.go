package examplePlugin

import (
	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

const name = "ExamplePlugin"

var log *logrus.Logger

type Example struct{}

func init() {
	loader.LoadFunctions = append(loader.LoadFunctions, Load)
}

func Load(logger *logrus.Logger) (plugins.Plugin, error) {
	log = logger
	log.Tracef("%s:\tLoad\n", name)

	return Example{}, nil
}

func (ex Example) Run() error {
	log.Tracef("%s:\tRun\n", name)
	return nil
}

func (ex Example) Name() string {
	return name
}
