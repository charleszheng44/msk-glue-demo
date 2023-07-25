package util

import (
	"msk-glue-demo/pkg/config"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/segmentio/kafka-go"
)

func CreateTopic(topic string) {
	kafkaServer := GetEnv(config.KafkaServersEnv)
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0

	admin, err := sarama.NewClusterAdmin(strings.Split(kafkaServer, ","), config)
	if err != nil {
		panic("Error creating cluster admin: " + err.Error())
	}
	defer func() {
		if err := admin.Close(); err != nil {
			panic("Error closing cluster admin: " + err.Error())
		}
	}()

	detail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	err = admin.CreateTopic(topic, detail, false)
	if err != nil {
		panic("Error creating topic: " + err.Error())
	}
}

func NewKafkaWriter() *kafka.Writer {
	kafkaServer := GetEnv(config.KafkaServersEnv)
	return kafka.NewWriter(
		kafka.WriterConfig{
			Brokers: strings.Split(kafkaServer, ","),
			Topic:   config.TopicName,
		},
	)
}

func NewKafkaReader() *kafka.Reader {
	kafkaServer := GetEnv(config.KafkaServersEnv)
	return kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: strings.Split(kafkaServer, ","),
			Topic:   config.TopicName,
		},
	)
}
