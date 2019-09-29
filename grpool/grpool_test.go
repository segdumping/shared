package grpool

import (
	"testing"
	"time"
)

func TestNewGRPool(t *testing.T) {
	grpool := NewGRPool(100, 50)
	defer grpool.Release()
	grpool.Concurrent(func(pool *GRPool) {
		for i := 0; i < 10; i++ {
			count := i
			grpool.WaitCount(1)
			pool.JobQueue <- func() {
				defer grpool.JobDone()
				t.Log(count)
			}
		}
	})

	grpool.WaitAll()

}

func TestConcurrent(t *testing.T) {
	Concurrent(func(pool *GRPool) {
		for i := 0; i < 10; i++ {
			count := i
			pool.JobQueue <- func() {
				t.Log(count)
			}
		}
	})

	timer := time.NewTicker(1 * time.Second)
	select {
	case <-timer.C:
	}

	t.Log("concurrent finish")
}

func TestConcurrentAutoWait(t *testing.T) {
	var step = 3
	var total = 10
	var start, end int
	ConcurrentAutoWait(total, step, func(pool *GRPool) {
		for end < total {
			if start > total-step {
				end = total
			} else {
				end = start + step
			}

			s, e := start, end
			pool.JobQueue <- func() {
				defer pool.JobDone()
				t.Log("start:", s, "end:", e)
			}

			start = end
		}
	})

	t.Log("concurrentautowait finish")
}

func TestConcurrentAddWait(t *testing.T) {
	ConcurrentAddWait(func(pool *GRPool) {
		for i := 0; i < 10; i++ {
			count := i
			pool.WaitCount(1)
			pool.JobQueue <- func() {
				defer pool.JobDone()
				t.Log(count)
			}
		}
	})

	t.Log("concurrentaddwait finish")
}

func TestGoroutineRun(t *testing.T) {
	GoroutineRun(func() {
		t.Log("TestGoroutineRun")
	})

	time.Sleep(time.Second)
}
