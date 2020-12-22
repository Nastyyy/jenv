package actions

type Action interface {
	DoAction() error
	GetHelp() string
}
