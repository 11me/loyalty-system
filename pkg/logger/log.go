package logger

import "github.com/sirupsen/logrus"

var rootLogger *logrus.Logger

type Fields = logrus.Fields

type Logger struct {
	*logrus.Entry
}

func InitLogger(lvl string) error {
	rootLogger = logrus.New()
	rootLogger.SetFormatter(&logrus.JSONFormatter{})
	logLvl, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLvl)
	return nil
}

func GetLogger() Logger {
	return Logger{logrus.NewEntry(rootLogger)}
}
