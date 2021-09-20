package curl

import (
	"strconv"
	"time"
)

type Seconds float64

func NewSeconds(value time.Duration) *Seconds {
	s := Seconds(value.Seconds())
	return &s
}

func (s *Seconds) Set(value string) error {
	v, err := parseFloat(value)
	*s = Seconds(v)
	return err
}

func (s Seconds) Get() interface{} { return time.Duration(s * 1e9) }
func (s Seconds) String() string   { return formatFloat(float64(s)) }
func (s Seconds) Type() string     { return "seconds" }

func secondsOption(name, short string, defval time.Duration, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewSeconds(defval)}
}

func parseFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}

func formatFloat(f float64) string {
	if f == 0 {
		return ""
	}
	return strconv.FormatFloat(f, 'g', -1, 64)
}

type Milliseconds int64

func NewMilliseconds(value time.Duration) *Milliseconds {
	ms := Milliseconds(value / time.Millisecond)
	return &ms
}

func (ms *Milliseconds) Set(value string) error {
	v, err := parseInt64(value)
	*ms = Milliseconds(v)
	return err
}

func (ms Milliseconds) Get() interface{} { return time.Duration(ms * 1e6) }
func (ms Milliseconds) String() string   { return formatInt64(int64(ms)) }
func (ms Milliseconds) Type() string     { return "milliseconds" }

func millisecondsOption(name, short string, defval time.Duration, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewMilliseconds(defval)}
}

type Speed string

func NewSpeed(value string) *Speed {
	s := Speed(value)
	return &s
}

func (s *Speed) Set(value string) error {
	*s = Speed(value)
	return nil
}

func (s Speed) Get() interface{} { return string(s) }
func (s Speed) String() string   { return string(s) }
func (s Speed) Type() string     { return "speed" }

func speedOption(name, short, defval, help string) Option {
	return Option{Name: name, Help: help, Short: short, Value: NewSpeed(defval)}
}
