package log

import (
	"log/slog"
	"os"
)

var StdLogger *slog.Logger
var FileLogger *slog.Logger

type LogMode uint8

const (
	UNDEFINE LogMode = 0
	STDOUT   LogMode = 1
	FILEOUT  LogMode = 2
)

var logMode LogMode

func Init(logMode_ LogMode) {
	logMode = logMode_
	switch logMode_ {
	case STDOUT:
		// 打开一个标准输出用于输出日志
		StdLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(StdLogger)
	case FILEOUT:
		// 打开一个文件用于写入日志
		logFile, err := os.OpenFile("log_file.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// 无法打开日志文件，进行处理，比如输出到标准错误
			panic(err)
		}
		// 创建一个新的日志处理器，输出到文件
		FileLogger = slog.New(slog.NewJSONHandler(logFile, nil))
		// 设置为默认的日志处理器
		slog.SetDefault(FileLogger)
	}
}

func Info(msg string, args ...any) {
	switch logMode {
	case STDOUT:
		StdLogger.Info(msg, args)
	case FILEOUT:
		FileLogger.Info(msg, args)
	default:
		slog.Info(msg, args)
	}
}
func Warn(msg string, args ...any) {
	switch logMode {
	case STDOUT:
		StdLogger.Warn(msg, args)
	case FILEOUT:
		FileLogger.Warn(msg, args)
	default:
		slog.Warn(msg, args)
	}
}
func Debug(msg string, args ...any) {
	switch logMode {
	case STDOUT:
		StdLogger.Debug(msg, args)
	case FILEOUT:
		FileLogger.Debug(msg, args)
	default:
		slog.Debug(msg, args)
	}
}
func Error(msg string, args ...any) {
	switch logMode {
	case STDOUT:
		StdLogger.Error(msg, args)
	case FILEOUT:
		FileLogger.Error(msg, args)
	default:
		slog.Error(msg, args)
	}
}
