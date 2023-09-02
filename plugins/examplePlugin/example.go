package examplePlugin

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/communication"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

const name = "ExamplePlugin"

var log *logrus.Logger

var rChannel *communication.Channels

type Example struct{}

func init() {
	loader.LoadFunctions = append(loader.LoadFunctions, Load)
}

func Load(logger *logrus.Logger, channels communication.Channels) (plugins.Plugin, error) {
	log = logger
	log.Tracef("%s:\tLoad\n", name)

	rChannel = &channels

	return Example{}, nil
}

func (ex Example) Run() error {
	log.Tracef("%s:\tRun\n", name)

	for {
		msg := communication.Mqtt{
			Msg:   "This is a Test from the Example Plugin",
			Topic: "sys/1",
		}

		rChannel.Messages <- msg

		time.Sleep(20 * time.Second)
	}

	return nil
}

func (ex Example) Name() string {
	return name
}
