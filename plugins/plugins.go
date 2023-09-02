package plugins

type Plugin interface {
	Run()
	Name() string
}
