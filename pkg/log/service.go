package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
	INDEX_LOG_LOGIN    = "log_login"
)

func InsertErrorLog(ctx context.Context, log *LogError) error {
	return nil
}

func InsertActivityLog(ctx context.Context, log *LogError) error {
	return nil
}

func InsertLoginLog(ctx context.Context, log *LogError) error {
	return nil
}

func LogruswriteError(id string, message string) error {
	logrusvar := logrus.New()
	logrusvar.SetReportCaller(true)
	logrusvar.SetOutput(os.Stdout)
	logrusvar.SetLevel(logrus.DebugLevel)
	logrusvar.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrusvar.WithFields(logrus.Fields{
		"id":    id,
		"error": message,
	}).Error("An error occured")
	return nil
}
