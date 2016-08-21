package rxGo

type justOp struct{
  value interface{}
}

func (j *justOp) Subscribe(o Observer){
  o.OnSubscribe(createEmptyDisposable())
  o.OnNext(j.value)
  o.OnComplete()
}
