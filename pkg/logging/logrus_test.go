package logging

import "testing"

func TestLogrusLogger(t *testing.T) {
	logger := Logrus()
	logger.Debug("test %s", "1")
	logger.Info("test %s", "1")
	logger.Warn("test %s", "1")
	logger.Error("test %s", "1")
}

