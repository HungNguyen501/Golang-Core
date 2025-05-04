package logging

import (
	"time"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogging() *zap.SugaredLogger {
	var logger *zap.Logger
	logger = newLocalLogger()
	zap.ReplaceGlobals(logger)
	return logger.Sugar()
}

func newLocalLogger() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     zapcore.OmitKey,
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "M",
		StacktraceKey: "S",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2000-01-01 00:00:00.000"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(colorable.NewColorableStdout()),
			getLogLevel("local"),
		),
		zap.AddCaller(),
		zap.Development(),
	)
}

// Returns appropriate log level based on environment
func getLogLevel(env string) zapcore.Level {
	switch env {
	case "prod", "production":
		return zapcore.InfoLevel
	default:
		return zapcore.DebugLevel
	}
}
