package kafkapr

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func (k *KafkaPresentation) consumeTrades(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.KafkaAddresses,
		Topic:   k.topicTrade,
		GroupID: k.KafkaGroupId,
	})
	logData := logrus.Fields{
		"brokers": k.KafkaAddresses,
		"topic":   k.topicTrade,
		"group":   k.KafkaGroupId,
		"reader":  r,
	}
	logrus.WithFields(logData).Debug("consuming trades topic has been started")

	// consuming is here ...

	return

}
