package task

import (
	"fmt"
	"testing"
	"time"
)

func TestAddFuncWithInterval(t *testing.T)  {
	AddFuncWithInterval(time.Second, func() {
		fmt.Println("task job")
	})

	timer := time.NewTimer(5 * time.Second)

	select {
	case <- timer.C:
	}
}

type job struct {

}

func (j job) Run()  {
	fmt.Println("test job")
}

func TestAddJobWithInterval(t *testing.T) {
	AddJobWithInterval(time.Second, job{})

	timer := time.NewTimer(5 * time.Second)

	select {
	case <- timer.C:
	}
}

func TestRemoveTask(t *testing.T) {
	taskID := AddJobWithInterval(time.Second, job{})

	timer := time.NewTimer(5 * time.Second)

	select {
	case <- timer.C:
		RemoveTask(taskID)
	}

	timer.Reset(5 * time.Second)
	select {
	case <- timer.C:
	}
}