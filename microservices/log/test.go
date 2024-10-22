package log

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	t.Skip("implementation incomplete")

	log := Logger()
	log.Info("data", zapcore.Field{
		Key:     "time",
		Type:    zapcore.TimeType,
		Integer: time.Now().UTC().UnixMicro(),
	}, zapcore.Field{
		Key:     "displacement",
		Type:    zapcore.Float64Type,
		Integer: 0,
	})
}
