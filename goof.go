package goof

import (
	"context"
	"sync"
	"sync/atomic"
)

// Task is a first result aware struct.
type Task[T any] struct {
	res    *T
	cond   *sync.Cond
	c      int32
	ctx    context.Context
	cancel func()
}

// New create a first result aware Task.
func New[T any]() *Task[T] {
	t := &Task[T]{
		cond: sync.NewCond(&sync.Mutex{}),
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	return t
}

// Go runs the result generating task.
func (t *Task[T]) Go(fn func(ctx context.Context) (T, error)) {
	atomic.AddInt32(&t.c, 1)
	go func() {
		defer t.cond.Signal()
		defer atomic.AddInt32(&t.c, -1)
		res, err := fn(t.ctx)
		if err != nil {
			return
		}
		t.cond.L.Lock()
		if t.res == nil {
			t.res = &res
		}
		t.cond.L.Unlock()
	}()
}

// First returns the first successful result.
func (t *Task[T]) First() (T, bool) {
	defer t.cancel()
	t.cond.L.Lock()
	for {
		if t.res != nil {
			res := *t.res
			t.cond.L.Unlock()
			return res, true
		}
		if atomic.LoadInt32(&t.c) == 0 {
			t.cond.L.Unlock()
			return *new(T), false
		}
		t.cond.Wait()
	}
}
