package examplePlugin

import (
	"fmt"

	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

const name = "ExamplePlugin"

type Example struct{}

func init() {
	fmt.Printf("%s:\tregistering..\n", name)
	loader.LoadFunctions = append(loader.LoadFunctions, Load)
}

func Load() (plugins.Plugin, error) {
	fmt.Printf("%s:\tLoad\n", name)

	return Example{}, nil
}

func (ex Example) Run() error {
	fmt.Printf("%s:\tRun\n", name)
	return nil
}

func (ex Example) Name() string {
	return name
}
