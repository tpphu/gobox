package provider

type Provider interface {
	Init()
	Run()
}

func NewProvider() Provider {
	p := provider{}
	return &p
}

type provider struct {
}

func (p *provider) Init() {
}

func (p *provider) Run() {
}
