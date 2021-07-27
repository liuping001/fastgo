// Author: coolliu
// Date: 2021/7/27

package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapLog struct {
	Logger lumberjack.Logger `json:"logger" yaml:"logger"`
	Level  zapcore.Level     `json:"level" yaml:"level"`
}

// 注意需要defer
// zapLog = NewZapLogger(nil)
// defer zapLog.Sync()
func NewZapLogger(conf *ZapLog) *zap.SugaredLogger {
	if conf == nil {
		ll := lumberjack.Logger{
			Filename:   os.Args[0] + ".log",
			MaxSize:    1024 * 1024 * 100,
			MaxBackups: 10,
		}
		conf = &ZapLog{Logger: ll, Level: zapcore.DebugLevel}
	}
	writeSyncer := zapcore.AddSync(&conf.Logger)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, conf.Level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}
