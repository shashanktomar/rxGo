package rxGo

type Subscriber interface{
  OnSubscribe(s *Subscription)
  OnNext(value interface{})
  OnError(err error)
  OnComplete()
}
