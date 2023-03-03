package logger

import "fmt"

type Stdout struct{}

func (s *Stdout) Log(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

var _ Logger = &Stdout{}
