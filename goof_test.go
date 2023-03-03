package goof

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestGoofAllSuccess(t *testing.T) {
	tasks := New[int]()
	tasks.Go(func(ctx context.Context) (int, error) {
		time.Sleep(time.Millisecond * 10)
		return 1, nil
	})
	tasks.Go(func(ctx context.Context) (int, error) {
		time.Sleep(time.Millisecond * 20)
		return 2, nil
	})
	res, ok := tasks.First()
	if !ok {
		t.Errorf("ok should be true")
		return
	}
	if res != 1 {
		t.Errorf("result should be 1")
	}
}

func TestGoofSuccess(t *testing.T) {
	tasks := New[int]()
	tasks.Go(func(ctx context.Context) (int, error) {
		time.Sleep(time.Millisecond * 10)
		return 1, nil
	})
	tasks.Go(func(ctx context.Context) (int, error) {
		return 2, errors.New("2")
	})
	res, ok := tasks.First()
	if !ok {
		t.Errorf("ok should be true")
		return
	}
	if res != 1 {
		t.Errorf("result should be 1")
	}
}

func TestGoofFailure(t *testing.T) {
	tasks := New[int]()
	tasks.Go(func(ctx context.Context) (int, error) {
		return 0, errors.New("1")
	})
	tasks.Go(func(ctx context.Context) (int, error) {
		return 0, errors.New("2")
	})
	_, ok := tasks.First()
	if ok {
		t.Errorf("ok should be false")
		return
	}
}
