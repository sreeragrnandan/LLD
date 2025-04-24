package schedular

import "time"

type Schedular interface {
	Schedule(task func()) error
}

type SimpleSchedular struct{}

func (s *SimpleSchedular) Schedule(task func()) error {
	go func() {
		time.Sleep(1 * time.Second)
		task()
	}()
	return nil
}
