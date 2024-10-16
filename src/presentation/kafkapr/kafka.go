package kafkapr

import (
	"context"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
	"gitlab.com/hamidteimouri/core-lens/domain"
)

const (
	TopicTrades = "trades"
)

type KafkaPresentation struct {
	lensService    *domain.LensService
	topicTrade     string
	KafkaAddresses []string
	KafkaGroupId   string
}

func NewKafkaPresentation(lensService *domain.LensService) *KafkaPresentation {
	return &KafkaPresentation{
		lensService:    lensService,
		KafkaAddresses: htenvier.Envs("KAFKA_ADDRESS"),
		topicTrade:     TopicTrades,
		KafkaGroupId:   htenvier.Env("KAFKA_GROUP_ID"),
	}
}

// Consume from Kafka
func (k *KafkaPresentation) Consume(ctx context.Context) {
	// Get an array of Kafka nodes
	logrus.WithFields(logrus.Fields{
		"kafka_addresses": k.KafkaAddresses,
		"group_id":        k.KafkaGroupId,
		"topics":          TopicTrades,
	}).Info("preparing to consume from kafka")

	go k.consumeTrades(ctx)
}
