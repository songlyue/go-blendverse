package initialize

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

func InitLogger(filePath, infoFileName, warnFileName, fileExt string) (*zap.Logger, error) {
	encoder := getEncoder()
	//判断日志等级 warn一下属于info
	infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zapcore.WarnLevel
	})
	// 判断日志等级 wanrn及以上属于warn
	warnLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.WarnLevel
	})

	infoWriter, err := getLogWriter(filePath+"/"+infoFileName, fileExt)
	if err != nil {
		return nil, err
	}
	warnWriter, err2 := getLogWriter(filePath+"/"+warnFileName, fileExt)
	if err2 != nil {
		return nil, err2
	}
	// 创建具体的logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoWriter, infoLevel),
		zapcore.NewCore(encoder, warnWriter, warnLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil

}

// 编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 或者可以选json格式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filePath, fileExt string) (zapcore.WriteSyncer, error) {
	writer, err := getWriter(filePath, fileExt)
	if err != nil {
		return nil, err
	}
	// 输出到控制台
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)), nil
}

// 指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去。
func getWriter(filename, fileExt string) (io.Writer, error) {
	hook, err := rotatelogs.New(
		filename+"_%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Minute),
	)
	if err != nil {
		return nil, err
	}
	return hook, nil
}
