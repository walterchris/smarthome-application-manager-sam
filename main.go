package main

import (
	"fmt"

	"github.com/walterchris/smarthome-application-manager-sam/pkg/config"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	_ "github.com/walterchris/smarthome-application-manager-sam/plugins/caldev"
	_ "github.com/walterchris/smarthome-application-manager-sam/plugins/examplePlugin"
)

func main() {
	fmt.Println("Started Main.")

	// Load Configuration
	config, err := config.Parse("./config.yaml")
	if err != nil {
		panic(err.Error())
	}

	if config != nil {
		fmt.Printf("%+v\n", config)
	}

	// Load Plugin
	for _, loadfunc := range loader.LoadFunctions {
		if loadfunc != nil {
			p, err := loadfunc()
			if err != nil || p == nil {
				panic(err.Error())
			}

			// Check if Plugin is in Config - if not, don't run it.
			found := false
			for _, plugin := range config.Plugins {
				found = false
				if plugin[p.Name()] != nil {
					found = true
					break
				}
			}
			if !found {
				continue
			}

			// Run the Plugins
			err = p.Run()
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
