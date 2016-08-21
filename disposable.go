package rxGo

import "errors"

type Disposable interface {
	Dispose()
	IsDisposed() bool
}

type EmptyDisposable struct{}

func (e *EmptyDisposable) Dispose() {}
func (e *EmptyDisposable) IsDisposed() bool {
	return true
}

func createEmptyDisposable() Disposable {
	return &EmptyDisposable{}
}

func validateDisposable(current Disposable, next Disposable) error {
	if current != nil {
		return errors.New("current disposable should be nil")
	}
	if next == nil {
		return errors.New("next disposable should not be nil")
	}
	return nil
}
