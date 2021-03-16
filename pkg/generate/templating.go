package generate

type templating interface {
	Name() string
	Template() string
}
