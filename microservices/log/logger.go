package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level = zapcore.DebugLevel

func Logger() *zap.Logger {
	distributedDBEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	syncedWriter := zapcore.AddSync(&Writer{})

	distributedDBCore := zapcore.NewCore(
		distributedDBEncoder,
		syncedWriter,
		level,
	)

	core := zapcore.NewTee(
		distributedDBCore,
		// add more cores
	)

	return zap.New(core)
}
