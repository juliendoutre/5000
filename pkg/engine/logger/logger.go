package logger

type Logger interface {
	Log(log string, args ...any)
}
