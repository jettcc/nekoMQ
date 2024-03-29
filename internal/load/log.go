package load

import (
	"sesm/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func Init(logPath string) {
	hook := lumberjack.Logger{
		Filename: logPath,
	}
	w := zapcore.AddSync(&hook)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		zap.InfoLevel,
	)

	logger = zap.New(core, zap.AddCaller())
	global.Sugar = logger.Sugar()
	return
}

func Sync() {
	logger.Sync()
}
