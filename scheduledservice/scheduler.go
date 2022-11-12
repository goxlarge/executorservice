package scheduledservice

import "time"

type Runnable interface {
	Run()
}

func Schedulejob(interval int64, runnable Runnable) chan bool {
	quit := make(chan bool)
	go func() {
		for {
			runnable.Run()
			select {
			case <-quit:
				return
			case <-time.After(time.Duration(interval) * time.Second):
				// repeats loop
			}
		}
	}()
	return quit
}

func SchedulejobWith(interval int64, runnable Runnable, last int64) <-chan time.Time {
	ch := time.After(time.Duration(last) * time.Second)
	go func() {
		for {
			runnable.Run()
			select {
			case <-ch:
				return
			case <-time.After(time.Duration(interval) * time.Second):
				// repeats loop
			}
		}
	}()
	return ch
}
