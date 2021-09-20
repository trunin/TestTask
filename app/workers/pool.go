package workers

import (
	"context"
)

type (
	PoolType chan struct{}
)

type getUrlCallable func(ctx context.Context, url string,
	outCh chan []byte, errCh chan error)

type Pool struct {
	numberOfWorkers int
	pool            PoolType
	callable        getUrlCallable
}

func (s *Pool) Stop() {
	for i := 0; i < s.numberOfWorkers; i++ {
		s.pool <- struct{}{}
	}
}

func (s *Pool) Do(ctx context.Context, url string,
	outCh chan []byte, errCh chan error) {
	s.pool <- struct{}{}
	go func() {
		s.callable(ctx, url, outCh, errCh)
		<-s.pool
	}()
}

func NewPool(numberOfWorkers int, fn getUrlCallable) *Pool {
	smf := Pool{numberOfWorkers, make(chan struct{}, numberOfWorkers), fn}
	return &smf
}
