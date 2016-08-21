package rxGo

type fromArrayOp struct {
	array []interface{}
}

func (j *fromArrayOp) Subscribe(o Observer) {
	o.OnSubscribe(createEmptyDisposable())
	o.OnComplete()
}
