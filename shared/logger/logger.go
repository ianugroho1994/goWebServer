package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

// Log is global logger
var Log *Logger

// Logger ...
type Logger struct {
	log *logrus.Logger
}

func initializeLogRotator(log *logrus.Logger, path string) error {
	if path == "" {
		path = "./logs/log"
	}
	//path += ".%Y%m%d%H%M"

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(path),
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)

	log.SetOutput(multiWriter)
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - {%customField%} - %msg% \n",
	})

	return nil
}

func NewLogger(path string) *Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	error := initializeLogRotator(log, path)
	if error != nil {
		fmt.Println(error)
		return nil
	}

	return &Logger{
		log: log,
	}
}

func (l *Logger) SetLogLevel(level logrus.Level) {
	l.log.SetLevel(level)
}

// Panic logs a message at level Panic on the standard logger
func (l *Logger) Panic(args ...interface{}) {
	if l.log.Level >= logrus.PanicLevel {
		l.Print(logrus.PanicLevel, args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger
func (l *Logger) Fatal(args ...interface{}) {
	if l.log.Level >= logrus.FatalLevel {
		l.Print(logrus.FatalLevel, args...)
	}
}

// Error logs a message at level Error on the standard logger
func (l *Logger) Error(args ...interface{}) {
	if l.log.Level >= logrus.ErrorLevel {
		l.Print(logrus.ErrorLevel, args...)
	}
}

// Warn logs a message at level Warn on the standard logger
func (l *Logger) Warn(args ...interface{}) {
	if l.log.Level >= logrus.WarnLevel {
		l.Print(logrus.WarnLevel, args...)
	}
}

// Info logs a message at level Info on the standard logger.
func (l *Logger) Info(args ...interface{}) {
	if l.log.Level >= logrus.InfoLevel {
		l.Print(logrus.InfoLevel, args...)
	}
}

// Debug logs a message at level Debug on the standard logger
func (l *Logger) Debug(args ...interface{}) {
	if l.log.Level >= logrus.DebugLevel {
		l.Print(logrus.DebugLevel, args...)
	}
}

// Trace logs a message at level Trace on the standard logger
func (l *Logger) Trace(args ...interface{}) {
	if l.log.Level >= logrus.TraceLevel {
		l.Print(logrus.TraceLevel, args...)
	}
}

func (l *Logger) Print(level logrus.Level, args ...interface{}) {
	entry := l.log.WithField("customField", getFileInfo(3))
	entry.Logln(level, args...)
}

func getFileInfo(skip int) string {
	counter, file, line, ok := runtime.Caller(skip)
	functionName := ""
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
		functionName = runtime.FuncForPC(counter).Name()
	}
	return fmt.Sprintf("%s() - %s:%d", functionName, file, line)
}
