package service

import "context"

type Runable interface {
	Init() error
	Run() error
	Shutdown(ctx context.Context) error

}
