package mqtt

import "github.com/walterchris/smarthome-application-manager-sam/pkg/communication"

type MQTT struct {
	Channels communication.Channels
}

func New() (*MQTT, error) {
	return &MQTT{
		Channels: communication.Channels{},
	}, nil
}

func (mqtt *MQTT) Run() {
	// Set up MQTT, populate topics (How do know what plugins need what topics?)
}
