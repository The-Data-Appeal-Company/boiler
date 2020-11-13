package logging

type NoopLogger struct {
}

func Noop() NoopLogger {
	return NoopLogger{}
}

func (n NoopLogger) Debug(s string, i ...interface{}) {
}

func (n NoopLogger) Info(s string, i ...interface{}) {
}

func (n NoopLogger) Warn(s string, i ...interface{}) {
}

func (n NoopLogger) Error(s string, i ...interface{}) {
}
