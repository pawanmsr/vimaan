package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logger(level zapcore.Level) *zap.Logger {
	distributedDBEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	syncedWriter := zapcore.AddSync(Writer{})

	// verify level

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
