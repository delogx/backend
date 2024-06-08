package services

type Service struct{}

type App struct {
	ID       uint
	HostName string
}

func (sc Service) FindOne() App {
	return App{}
}
