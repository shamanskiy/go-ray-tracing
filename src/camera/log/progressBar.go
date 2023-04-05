package log

import "github.com/schollz/progressbar/v3"

const (
	UNSET_MAX      = 1
	NO_DESCRIPTION = ""
)

type ProgressUpdate struct {
	Max int
}

func NewProgressBar() (progressChan chan<- ProgressUpdate) {
	bar := progressbar.Default(UNSET_MAX, NO_DESCRIPTION)

	c := make(chan ProgressUpdate)
	go func() {
		for msg := range c {
			bar.ChangeMax(msg.Max)
			bar.Add(1)
		}
	}()

	return c
}
