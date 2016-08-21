package rxGo

type Subscription interface{
    Request(n int64)
    Cancel()
}
