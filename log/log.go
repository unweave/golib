package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

var (
	Logger Log
)

type Log struct {
	pretty *zap.SugaredLogger
	fast   *zap.Logger
}

func init() {
	Logger = NewLogger()
}

func SetOutput(output io.Writer) {
	zapcore.AddSync(output)
}

func NewLogger() Log {
	opt := zap.AddCallerSkip(1)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	logger, err := config.Build(opt)

	if err != nil {
		panic(err)
	}
	return Log{
		pretty: logger.Sugar(),
		fast:   logger,
	}
}

func Info(msg string, args ...interface{}) {
	Logger.pretty.Infof(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	Logger.pretty.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	Logger.pretty.Warnf(msg, args...)
}

func Debug(msg interface{}) {
	Logger.pretty.Debug(msg)
}

func Debugf(msg string, args ...interface{}) {
	Logger.pretty.Debugf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	Logger.pretty.Error(msg, args)
}

func Errorf(msg string, args ...interface{}) {
	Logger.pretty.Errorf(msg, args)
}

func Fatal(msg string, args ...interface{}) {
	Logger.pretty.Fatalf(msg, args)
}

func Fatalf(msg string, args ...interface{}) {
	Logger.pretty.Fatalf(msg, args)
}
