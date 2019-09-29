package grpool

import "github.com/ivpusic/grpool"

var pool *GRPool

type GRPool struct {
	*grpool.Pool
}

func init() {
	pool = NewGRPool(100, 50)
}

func NewGRPool(workersNum, jobQueueLen int) *GRPool {
	return &GRPool{grpool.NewPool(workersNum, jobQueueLen)}
}

func (e *GRPool) Concurrent(jobs func(pool *GRPool)) {
	jobs(e)
}

func (e *GRPool) ConcurrentAddWait(jobs func(pool *GRPool)) {
	jobs(e)
	pool.WaitAll()
}

func (e *GRPool) ConcurrentAutoWait(total, step int, jobs func(pool *GRPool)) {
	pool.WaitCount(ChunkCount(total, step))
	jobs(pool)
	pool.WaitAll()
}

func Concurrent(jobs func(pool *GRPool)) {
	pool.Concurrent(jobs)
}

func ConcurrentAddWait(jobs func(pool *GRPool)) {
	pool.ConcurrentAddWait(jobs)
}

func ConcurrentAutoWait(total, step int, jobs func(pool *GRPool)) {
	pool.ConcurrentAutoWait(total, step, jobs)
}

func GoroutineRun(job func()) {
	pool.JobQueue <- job
}

func ChunkCount(total, step int) int {
	if total <= step {
		return 1
	}

	if step == 0 {
		return total
	}

	count := total / step
	if total%step > 0 {
		count++
	}

	return count
}
