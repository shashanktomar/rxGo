package rxGo

type Observable struct {
	delegate ObservableSource
}

func (o *Observable) Subscribe(observer Observer) {
	o.delegate.Subscribe(observer)
}

func Just(value interface{}) *Observable {
	return &Observable{&justOp{value}}
}

func (o *Observable) Map(mapper func(from interface{}) (interface{}, error)) *Observable {
	return &Observable{&mapOp{o, mapper}}
}
