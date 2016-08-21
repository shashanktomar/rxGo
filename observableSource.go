package rxGo

type ObservableSource interface {
	Subscribe(observer Observer)
}
