package logx

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() error {
	encoder := getEncoder()

	level := new(zapcore.Level)
	if err := level.UnmarshalText([]byte(viper.GetString("service.log.level"))); err != nil {
		return err
	}

	var core zapcore.Core
	if viper.GetString("service.mode") == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, getLogWriter(filepath.Join(viper.GetString("service.log.path"), "info.log"), viper.GetInt("service.log.rotateSize"), viper.GetInt("service.log.backups"), viper.GetInt("service.log.rotateDays")), level),
			zapcore.NewCore(encoder, getLogWriter(filepath.Join(viper.GetString("service.log.path"), "error.log"), viper.GetInt("service.log.rotateSize"), viper.GetInt("service.log.backups"), viper.GetInt("service.log.rotateDays")), zapcore.ErrorLevel),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, getLogWriter(filepath.Join(viper.GetString("service.log.path"), "info.log"), viper.GetInt("service.log.rotateSize"), viper.GetInt("service.log.backups"), viper.GetInt("service.log.rotateDays")), level),
			zapcore.NewCore(encoder, getLogWriter(filepath.Join(viper.GetString("service.log.path"), "error.log"), viper.GetInt("service.log.rotateSize"), viper.GetInt("service.log.backups"), viper.GetInt("service.log.rotateDays")), zapcore.ErrorLevel),
		)
	}
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return nil
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
