package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)

	Log.SetOutput(io.Discard)

	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic("failed to create log directory: " + err.Error())
	}

	logFileName := fmt.Sprintf("logs/app-%s.log", time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	consoleHook := &writerHook{
		Writer: os.Stdout,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
			FullTimestamp:   true,
		},
		LogLevels: logrus.AllLevels,
	}

	fileHook := &writerHook{
		Writer: logFile,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors:   true,
			FullTimestamp:   true,
		},
		LogLevels: logrus.AllLevels,
	}

	Log.AddHook(consoleHook)
	Log.AddHook(fileHook)
}

type writerHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
	LogLevels []logrus.Level
}

func (h *writerHook) Fire(entry *logrus.Entry) error {
	msg, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.Writer.Write(msg)
	return err
}

func (h *writerHook) Levels() []logrus.Level {
	return h.LogLevels
}
