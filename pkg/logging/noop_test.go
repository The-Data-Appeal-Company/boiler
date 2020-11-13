package logging

import "testing"

func TestNoopLogger(t *testing.T) {
	logger := Noop()
	logger.Debug("test %s", "1")
	logger.Info("test %s", "1")
	logger.Warn("test %s", "1")
	logger.Error("test %s", "1")
}
