package log

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"go.uber.org/zap"
	"os"
)

var logger *zap.Logger

func init() {
	logger = BuildLogger("example.log")
}

func BuildLogger(fileName string) *zap.Logger {
	// TODO use viper
	logDir := "./"
	if os.Getenv("EXAMPLE_LOG_PATH") != "" {
		logDir = os.Getenv("EXAMPLE_LOG_PATH")
	}

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logDir+"/"+fileName,
		MaxSize:    250, // megabytes
		MaxBackups: 54,
		MaxAge:     28, // days
	})

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.CallerKey = "file"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		writer,
		zap.InfoLevel,
	)

	zapLogger := zap.New(core, zap.AddCaller())

	return zapLogger
}

func Logger() *zap.Logger {
	return logger
}

func Info(msg string, fields ...zapcore.Field)  {
	logger.Info(msg, fields...)
}