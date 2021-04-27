package logger

import (
	"go-ant-design-pro/constant"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var zapSugaredLogger *zap.SugaredLogger
var zapLogger *zap.Logger

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	level := getLogLevel()
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// Log, _ = zap.NewProduction()
	zapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapSugaredLogger = zapLogger.Sugar()
	zapSugaredLogger.Sync()
}

func PluginInGin(router *gin.Engine) {
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(zapLogger, true))
}

/*
	info log level by default
*/
func getLogLevel() zapcore.Level {

	var logLevel string
	if logLevel = os.Getenv("LOG_LEVEL"); len(logLevel) != 0 {
		logLevel = constant.DEFAULT_LOG_LEVEL
	}

	defaultLogLevel := zapcore.DebugLevel

	switch logLevel {
	case "debug", "DEBUG":
		defaultLogLevel = zapcore.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		defaultLogLevel = zapcore.InfoLevel
	case "warn", "WARN":
		defaultLogLevel = zapcore.WarnLevel
	case "error", "ERROR":
		defaultLogLevel = zapcore.ErrorLevel
	case "dpanic", "DPANIC":
		defaultLogLevel = zapcore.DPanicLevel
	case "panic", "PANIC":
		defaultLogLevel = zapcore.PanicLevel
	case "fatal", "FATAL":
		defaultLogLevel = zapcore.FatalLevel
	}

	return defaultLogLevel
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder
	// zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	var filename string
	if filename = os.Getenv("LOG_FILENAME"); len(filename) == 0 {
		filename = constant.DEFAULT_LOG_FILENAME
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(args ...interface{}) {
	zapSugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapSugaredLogger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	zapSugaredLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapSugaredLogger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	zapSugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapSugaredLogger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	zapSugaredLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapSugaredLogger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	zapSugaredLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	zapSugaredLogger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Panicw(msg, keysAndValues...)
}
