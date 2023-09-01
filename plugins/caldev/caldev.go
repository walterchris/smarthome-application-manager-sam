package caldev

import (
	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/communication"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

const name = "CalDev-Plugin"

var log *logrus.Logger

type Caldev struct{}

func init() {
	loader.LoadFunctions = append(loader.LoadFunctions, Load)
}

func Load(logger *logrus.Logger, channels communication.Channels) (plugins.Plugin, error) {
	log = logger
	log.Tracef("%s:\tLoad\n", name)

	return Caldev{}, nil
}

func (cd Caldev) Run() error {
	log.Tracef("%s:\tRun\n", name)

	return nil
}

func (cd Caldev) Name() string {
	return name
}
