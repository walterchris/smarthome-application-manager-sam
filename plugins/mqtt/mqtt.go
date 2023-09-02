package mqtt

import (
	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/communication"
)

var log *logrus.Logger

type MQTT struct {
	Channels communication.Channels
}

func New(logger *logrus.Logger) (*MQTT, error) {
	log = logger

	return &MQTT{
		Channels: communication.Channels{
			Messages: make(chan communication.Mqtt),
		},
	}, nil
}

func (mqtt *MQTT) Run() {
	// We only need a client to post to our MQTT broker
	for {
		mqttMsg := <-mqtt.Channels.Messages
		log.Printf("MQTT Message '%s' on '%s'\n", mqttMsg.Msg, mqttMsg.Topic)
	}
}
