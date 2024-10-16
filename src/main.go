package main

import (
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
	"gitlab.com/hamidteimouri/core-lens/di"
	_ "gitlab.com/hamidteimouri/core-lens/pkg"
	"gitlab.com/hamidteimouri/core-lens/presentation/httppr"
	"gitlab.com/hamidteimouri/core-lens/presentation/kafkapr"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	checkEnvs()

	// Kafka Presentation
	if htenvier.EnvOrDefault("CONSUME_FROM_KAFKA", "true") == "true" {
		kpr := kafkapr.NewKafkaPresentation(di.DomainLensService())
		kpr.Consume(context.Background())
	} else {
		logrus.Info("consuming from Kafka has been disabled")
	}

	httppr.Start()

	//wait for ctrl+c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
	<-ch
}

func checkEnvs() {
	/* General */
	if htenvier.Env("APP_NAME") == "" {
		logrus.Panic("APP_NAME is required for log service name")
	}

	if htenvier.Env("LOG_LEVEL") == "" {
		logrus.Panic("LOG_LEVEL is required")
	}

	if htenvier.Env("HTTP_ADDRESS") == "" {
		logrus.Panic("HTTP_ADDRESS is required")
	}

	/* Kafka */
	if htenvier.Env("KAFKA_ADDRESS") == "" {
		logrus.Panic("KAFKA_ADDRESS is required")
	}
	if htenvier.Env("KAFKA_GROUP_ID") == "" {
		logrus.Panic("KAFKA_GROUP_ID is required")
	}
}
