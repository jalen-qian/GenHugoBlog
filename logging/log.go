package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var GLog *zap.Logger

func NewLogger(lvl string) (*zap.Logger, error) {
	logConf := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	// 更改日志记录的时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logConf.EncoderConfig = encoderConfig
	logConf.Encoding = "console"
	if err := logConf.Level.UnmarshalText([]byte(lvl)); err != nil {
		return nil, err
	}

	return logConf.Build(zap.WithCaller(true))
}
