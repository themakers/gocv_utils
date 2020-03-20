package debouncer

import (
	"context"
	"sync"
	"time"
)

type Debouncer struct {
	threshold time.Duration
	cancel    func()
	lock      sync.Mutex
}

func New(threshold time.Duration) *Debouncer {
	tt := &Debouncer{
		threshold: threshold,
	}
	return tt
}

func (deb *Debouncer) timer(ctx context.Context, fn func()) {
	select {
	case <-ctx.Done():
	case <-time.After(deb.threshold):
		fn()
	}
}

func (deb *Debouncer) Trigger(fn func()) {
	deb.lock.Lock()
	defer deb.lock.Unlock()

	if deb.cancel != nil {
		deb.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	deb.cancel = cancel

	go deb.timer(ctx, fn)
}
