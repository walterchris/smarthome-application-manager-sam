package caldev

import (
	"fmt"

	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

const name = "CalDev-Plugin"

type Caldev struct{}

func init() {
	fmt.Printf("%s:\tregistering..\n", name)
	loader.LoadFunctions = append(loader.LoadFunctions, Load)
}

func Load() (plugins.Plugin, error) {
	fmt.Printf("%s:\tLoad\n", name)

	return Caldev{}, nil
}

func (cd Caldev) Run() error {
	fmt.Printf("%s:\tRun\n", name)

	return nil
}

func (cd Caldev) Name() string {
	return name
}
