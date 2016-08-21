package rxGo

type Observer interface{
  OnSubscribe(d Disposable)
  OnNext(value interface{})
  OnError(err error)
  OnComplete()
}

type ObserverOne struct{
}

func (o *ObserverOne) OnSubscribe(d Disposable){}
func (o *ObserverOne) OnNext(value interface{}){
  println(value)
}
func (o *ObserverOne) OnError(err error){}
func (o *ObserverOne) OnComplete(){}
