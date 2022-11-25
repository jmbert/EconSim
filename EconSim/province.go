package sim

type Province struct {
	pop       *Pop
	factories []*Factory
}

func (p *Province) ProvinceSetup() {
	p.factories = append(p.factories, NewFactory(0))
	p.factories = append(p.factories, NewFactory(1))
	p.factories = append(p.factories, NewFactory(2))
	var pop = NewPop(0)
	p.pop = &pop
}
