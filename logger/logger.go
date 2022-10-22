package logger

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env       string
	SentryUrl string
	Version   string
}

func InitLogger(loggerConfig Config) {
	if loggerConfig.Env == "prod" {
		timeFormatLayout := "2006-01-02T15:04:05.000Z"
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "log_level",
			},
			TimestampFormat: timeFormatLayout,
		})
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	// Only init sentry if its url is available.
	if loggerConfig.SentryUrl != "" {
		logrus.Info("Setting Sentry Hook.")
		loadSentryConfig(loggerConfig)
	} else {
		logrus.Warn("Skipping Setting up Sentry Hook due to no SentryUrl is passed.")
	}
}

func loadSentryConfig(config Config) error {
	tags := map[string]string{
		"environment": config.Env,
		"version":     config.Version,
	}
	hook, err := logrus_sentry.NewWithTagsSentryHook(config.SentryUrl, tags, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})

	if err != nil {
		logrus.Error("Init sentry failed", err.Error())
		return err
	}
	hook.StacktraceConfiguration.Enable = true
	hook.StacktraceConfiguration.IncludeErrorBreadcrumb = true

	logrus.AddHook(hook)
	logrus.Info("Sentry initilization completed")
	return nil
}
