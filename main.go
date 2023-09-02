package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/config"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	_ "github.com/walterchris/smarthome-application-manager-sam/plugins/caldev"
	_ "github.com/walterchris/smarthome-application-manager-sam/plugins/deye600"
	_ "github.com/walterchris/smarthome-application-manager-sam/plugins/examplePlugin"
	"github.com/walterchris/smarthome-application-manager-sam/plugins/mqtt"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.TraceLevel)
}

func main() {
	log.Tracef("Starting main")

	// Create signals channel to run server until interrupted
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	// Load Configuration
	config, err := config.Parse("./config.yaml")
	if err != nil {
		log.Errorf("Unable to parse config file: %v", err)
		log.Errorf("Continuing with empty config resulting in loading no plugins")
	}

	if config != nil {
		log.Tracef("%+v\n", config)
	}

	// Load MQTT Communication
	mqtt, err := mqtt.New(log)
	if err != nil {
		log.Errorf("mqtt.New() = '%v'", err)
		return
	}

	// Load Plugin
	for _, loadfunc := range loader.LoadFunctions {
		if loadfunc != nil {
			p, err := loadfunc(log, mqtt.Channels)
			if err != nil || p == nil {
				log.Errorf("Excuting loading function failed with '%v' or was nil", err)
			}

			// Check if Plugin is in Config - if not, don't run it.
			found := false
			for _, plugin := range config.Plugins {
				found = false
				if plugin[p.Name()] != nil {
					log.Debugf("Found config for plugin '%s'", p.Name())
					found = true
					break
				}
			}
			if !found {
				continue
			}

			// Run the Plugins
			go p.Run()
		}
	}

	// Run MQTT
	go mqtt.Run()

	<-done
}
