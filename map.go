package rxGo

import "errors"

type mapOp struct {
	source *Observable
	mapper func(from interface{}) (interface{}, error)
}

func (m *mapOp) Subscribe(o Observer) {
	m.source.Subscribe(&mapper{observer: o, mapper: m.mapper})
}

type mapper struct {
	observer     Observer
	mapper       func(from interface{}) (interface{}, error)
	subscription Disposable
	completed    bool
}

func (m *mapper) OnSubscribe(d Disposable) {
	if validateDisposable(m.subscription, d) == nil {
		m.subscription = d
		m.observer.OnSubscribe(m)
	}
}

func (m *mapper) OnNext(value interface{}) {
	if m.completed {
		return
	}
	result, err := m.mapper(value)
	if err != nil {
		m.OnError(err)
		return
	}
	if result == nil {
		m.OnError(errors.New("map function generated a nil value"))
		return
	}
	m.observer.OnNext(result)
}

func (m *mapper) OnError(err error) {
	if m.completed {
		panic("Should not receive this call if after onComplete is called")
	}
	m.completed = true
	m.observer.OnError(err)
}

func (m *mapper) OnComplete() {
	if m.completed {
		return
	}
	m.completed = true
	m.observer.OnComplete()
}

func (m *mapper) Dispose() {
	m.subscription.Dispose()
}
func (m *mapper) IsDisposed() bool {
	return m.subscription.IsDisposed()
}
