package logger

type NoOp struct{}

func (n *NoOp) Log(format string, args ...any) {}

var _ Logger = &NoOp{}
