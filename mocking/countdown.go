package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
	write          = "write"
	sleep          = "sleep"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(writer io.Writer, s *SpyCountdownOperations) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		s.Write([]byte{})
		s.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
