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
	for _, chanel := range l.chArr {
		go func(c <-chan T) {
			for {
				select {
				case el := <-c:
					l.out <- el
				case <-ctx.Done():
					return
				}
			}
		}(chanel)
	}

	return l.out
}
