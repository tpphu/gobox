package service

type Runable interface {
	Init()
	Run()
	Shutdown()
}
