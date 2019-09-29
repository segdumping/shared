package task

import (
	"github.com/robfig/cron/v3"
	"time"
)

var c *cron.Cron

func init() {
	c = cron.New()
	go c.Run()
}

type schedule time.Duration

func (s schedule) Next(t time.Time) time.Time {
	return t.Add(time.Duration(s))
}

func AddFuncWithInterval(interval time.Duration, job func()) cron.EntryID {
	return c.Schedule(schedule(interval), cron.FuncJob(job))
}

func AddJobWithInterval(interval time.Duration, job cron.Job) cron.EntryID {
	return c.Schedule(schedule(interval), job)
}

func RemoveTask(id cron.EntryID) {
	c.Remove(id)
}
