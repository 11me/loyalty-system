package logger

import "github.com/sirupsen/logrus"

var rootLogger *logrus.Logger

type Fields = logrus.Fields

type Logger struct {
	*logrus.Entry
}

func InitLogger() {
	rootLogger = logrus.New()
	rootLogger.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

func GetLogger() Logger {
	return Logger{logrus.NewEntry(rootLogger)}
}
