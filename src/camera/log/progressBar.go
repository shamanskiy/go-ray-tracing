package log

import "github.com/schollz/progressbar/v3"

const UNSET_MAX = 1

type ProgressUpdate struct {
	Max int
}

func ProgressBar(processName string) (progressChan chan<- ProgressUpdate) {
	bar := progressbar.Default(UNSET_MAX, processName)

	c := make(chan ProgressUpdate)
	go func() {
		for msg := range c {
			bar.ChangeMax(msg.Max)
			bar.Add(1)
		}
	}()

	return c
}
