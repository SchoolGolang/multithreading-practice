package listener

import (
	"context"
)

type Listener[T any] interface {
	AddChan(<-chan T)
	Listen(context.Context) <-chan T
}

type GenericListener[T any] struct {
	chArr []<-chan T
	out   chan T
}

func NewListener[T any]() GenericListener[T] {
	return GenericListener[T]{
		chArr: make([]<-chan T, 0, 4),
		out:   make(chan T, 2),
	}
}

func (l *GenericListener[T]) AddChan(ch <-chan T) {
	l.chArr = append(l.chArr, ch)
}

func (l *GenericListener[T]) Listen(ctx context.Context) <-chan T {
	//TODO: implement listen functionality

	for i:=range l.chArr{
		go func(ch <- chan T ) {
			select {
			case Rec := <-ch:
			l.out <-Rec
			case <-ctx.Done():
				return
			}
		}(l.chArr[i])
	}
	return l.out
}
