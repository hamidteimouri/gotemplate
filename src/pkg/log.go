package pkg

import (
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logLevel := htenvier.EnvOrDefault("LOG_LEVEL", "info")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("LOG_LEVEL is not valid")
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: !htenvier.IsProduction(),
		TimestampFormat: func() string {
			if !htenvier.IsProduction() {
				return "2006-01-02 15:04:05"
			}
			return ""
		}(),
	})
}
