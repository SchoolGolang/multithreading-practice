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
	for _, ch := range l.chArr {
		go func(ch <-chan T) {
			for {
				select {
				case <-ctx.Done():
					return
				case data := <-ch:
					l.out <- data
				}
			}
		}(ch)
	} // Мені це не дуже подоається через те, що ми можемо створити забагато рутин.

	return l.out
}
