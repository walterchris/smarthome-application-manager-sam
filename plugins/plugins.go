package plugins

type Plugin interface {
	Run() error
	Name() string
}
