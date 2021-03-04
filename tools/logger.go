package tools

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rexwsd/blog/config"
	"github.com/rexwsd/blog/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

var Logger *zap.Logger
var LogerName string

func initLogger() {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 14:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	var infoWriter io.Writer
	var warnWriter io.Writer
	infoWriter = getWriter(LogerName)
	warnWriter = getWriter(LogerName)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), errorLevel),
	)
	Logger = zap.New(core, zap.AddCaller())
}

func GetLogger(logName string) *zap.Logger {
	LogerName = logName
	initLogger()
	return Logger
}
func getWriter(filename string) io.Writer {
	var infoPath string
	switch config.Env {
	case "dev":
		infoPath = ProjectPath() + "/" + setting.InfoPath + "/" + time.Now().Format("20060102") + "/" + filename + ".log"
		break
	case "production":
		infoPath = setting.InfoPath + "/" + time.Now().Format("20060102") + "/" + filename + ".log"
		break
	}
	hook, err := rotatelogs.New(
		infoPath,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
