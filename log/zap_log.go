// Author: coolliu
// Date: 2021/7/27

package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 注意需要defer
// zapLog = NewZapLogger(nil)
// defer zapLog.Sync()
func NewZapLogger(conf *lumberjack.Logger) *zap.SugaredLogger {
	if conf == nil {
		conf = &lumberjack.Logger{
			Filename:   os.Args[0] + ".log",
			MaxSize:    1024 * 1024 * 100,
			MaxBackups: 10,
		}
	}
	writeSyncer := zapcore.AddSync(conf)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}
