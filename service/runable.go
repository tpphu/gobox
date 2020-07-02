package service

type Runable interface {
	Init() error
	Run() error
	Shutdown() error
}
