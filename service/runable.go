package service

import "context"

type Runable interface {
	Init()
	Run() error
	Shutdown(ctx context.Context) error
}
