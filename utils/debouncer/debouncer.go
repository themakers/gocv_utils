package debouncer

import (
	"context"
	"sync"
	"time"
)

type Debouncer struct {
	threshold time.Duration
	cancels   map[string]func()
	lock      sync.Mutex
}

func New(threshold time.Duration) *Debouncer {
	tt := &Debouncer{
		threshold: threshold,
		cancels:   make(map[string]func()),
	}
	return tt
}

func (tt *Debouncer) timer(id string, ctx context.Context, fn func()) {
	select {
	case <-time.After(tt.threshold):
		(func() {
			// FIXME What to do with panics in fn callback
			//defer tools.Recover(func(rec interface{}, strace string) {
			//	log.Println("PANIC Debouncer::timer", id, rec, strace)
			//})

			tt.lock.Lock()
			defer tt.lock.Unlock()
			delete(tt.cancels, id)

			fn()
		})()
	case <-ctx.Done():
	}
}

func (tt *Debouncer) Trigger(id string, fn func()) {
	tt.lock.Lock()
	defer tt.lock.Unlock()

	if cancel, ok := tt.cancels[id]; ok {
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	tt.cancels[id] = cancel
	go tt.timer(id, ctx, fn)
}
