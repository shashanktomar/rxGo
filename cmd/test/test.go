package main

import (
	"github.com/rxGo"
)

type sub struct {
	rxGo.ObserverOne
}

func (s *sub) OnNext(value interface{}) {
	println(value.(int))
}

func main() {
	rxGo.Just(3).
		Map(func(from interface{}) (interface{}, error) {
			i := from.(int)
			return i * 3, nil
		}).
		Subscribe(&sub{})
}
