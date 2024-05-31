package zap

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	log "github.com/hexiaopi/data-space/pkg/logger"
)

type zapLogger struct {
	zapLogger *zap.Logger
}

type Config struct {
	FileName  string
	LogLevel  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
	Encoding  string
	Env       string
}

var Std log.Logger

func init() {
	Std = New(&Config{LogLevel: "info", Encoding: "console", Env: "dev"})
}

func New(conf *Config) log.Logger {
	lv := conf.LogLevel
	var level zapcore.Level
	switch strings.ToLower(lv) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	hook := lumberjack.Logger{
		Filename:   conf.FileName,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackup,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}

	var encoder zapcore.Encoder
	if conf.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "Logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		})
	} else {
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout)),
		level,
	)
	if conf.Env != "prod" {
		return &zapLogger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))}
	}
	return &zapLogger{zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02 15:04:05"))
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// // WithValue Adds a field to the specified context
// func (l *Logger) WithValue(ctx context.Context, fields ...zapcore.Field) context.Context {
// 	if c, ok := ctx.(*gin.Context); ok {
// 		ctx = c.Request.Context()
// 		c.Request = c.Request.WithContext(context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields...)))
// 		return c
// 	}
// 	return context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields...))
// }

// // WithContext Returns a zap instance from the specified context
//
//	func (l *Logger) WithContext(ctx context.Context) *Logger {
//		if c, ok := ctx.(*gin.Context); ok {
//			ctx = c.Request.Context()
//		}
//		zl := ctx.Value(ctxLoggerKey)
//		ctxLogger, ok := zl.(*zap.Logger)
//		if ok {
//			return &Logger{ctxLogger}
//		}
//		return l
//	}
func (l *zapLogger) Debug(msg string, fields ...log.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, args...)
}

func (l *zapLogger) Info(msg string, fields ...log.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.zapLogger.Sugar().Infof(format, args...)
}

func (l *zapLogger) Warn(msg string, fields ...log.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, args...)
}

func (l *zapLogger) Error(msg string, fields ...log.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, args...)
}

func (l *zapLogger) Panic(msg string, fields ...log.Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, args...)
}

func (l *zapLogger) Fatal(msg string, fields ...log.Field) {
	l.zapLogger.Fatal(msg, fields...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, args...)
}
