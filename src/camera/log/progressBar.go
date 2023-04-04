package log

import "github.com/schollz/progressbar/v3"

func ProgressBar(maxProgressValue int64, processName string) (progressChan chan<- int) {
	bar := progressbar.Default(maxProgressValue, processName)

	c := make(chan int)
	go func() {
		for progress := range c {
			bar.Set(progress)
		}
	}()

	return c
}
